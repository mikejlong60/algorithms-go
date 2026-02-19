package skiena_4

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

type Point struct{ X, Y float64 }

func dot(a, b Point) float64       { return a.X*b.X + a.Y*b.Y }
func sub(a, b Point) Point         { return Point{a.X - b.X, a.Y - b.Y} }
func add(a, b Point) Point         { return Point{a.X + b.X, a.Y + b.Y} }
func mul(a Point, s float64) Point { return Point{a.X * s, a.Y * s} }

func norm(a Point) float64 { return math.Hypot(a.X, a.Y) }
func unit(a Point) Point {
	n := norm(a)
	return Point{a.X / n, a.Y / n}
}
func perp(u Point) Point { // rotate +90°
	return Point{-u.Y, u.X}
}

type event struct {
	y     float64
	delta int // +1 at start, -1 at end
}

// MaxPenetrationLine returns the direction u (unit), the normal w (unit), the offset c,
// and the maximum number of boundary intersections bestK.
//
// The maximizing line is: dot(w, x) = c
// A point on the line is p0 = w * c (since w is unit).
func MaxPenetrationLine(poly []Point) (bestU Point, bestW Point, bestC float64, bestK int) {
	n := len(poly)
	bestK = -1

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v := sub(poly[j], poly[i])
			if norm(v) == 0 {
				continue
			}
			u := unit(v)

			k, c, ok := maxPenetrationsAndOffset(poly, u)
			if !ok {
				continue
			}
			if k > bestK {
				bestK = k
				bestU = u
				bestW = perp(u) // unit, since u is unit
				bestC = c
			}
		}
	}

	// If polygon is degenerate or candidates didn't run, bestK may stay -1.
	// For a valid simple polygon, you should get bestK >= 2.
	return
}

// For a fixed direction u, compute:
// - bestK: maximum number of boundary intersections among lines parallel to u
// - bestC: offset of a line achieving bestK (in dot(w,x)=c form), choosing an interior offset (mid-interval).
//
// Returns ok=false if it couldn't form a meaningful interval (e.g., no non-parallel edges).
func maxPenetrationsAndOffset(poly []Point, u Point) (bestK int, bestC float64, ok bool) {
	w := perp(u)

	events := make([]event, 0, 2*len(poly))
	const eps = 1e-12
	n := len(poly)

	for i := 0; i < n; i++ {
		a := poly[i]
		b := poly[(i+1)%n]

		ya := dot(a, w)
		yb := dot(b, w)

		if math.Abs(ya-yb) <= eps {
			// Edge parallel to u => it won't be crossed by generic lines parallel to u.
			continue
		}

		ylo, yhi := ya, yb
		if ylo > yhi {
			ylo, yhi = yhi, ylo
		}

		// Half-open interval [ylo, yhi):
		events = append(events, event{y: ylo, delta: +1})
		events = append(events, event{y: yhi, delta: -1})
	}

	if len(events) == 0 {
		return 0, 0, false
	}

	// Sort by y, and for equal y process starts (+1) BEFORE ends (-1)
	// so [ylo, yhi) is counted correctly.
	sort.Slice(events, func(i, j int) bool {
		if events[i].y == events[j].y {
			return events[i].delta > events[j].delta // +1 first
		}
		return events[i].y < events[j].y
	})

	bestK = 0
	bestC = events[0].y
	cur := 0

	// Sweep grouped by y; after applying all deltas at y, cur applies on [y, nextY)
	idx := 0
	for idx < len(events) {
		y := events[idx].y

		// apply all deltas at this y
		for idx < len(events) && events[idx].y == y {
			cur += events[idx].delta
			idx++
		}

		if idx >= len(events) {
			break
		}
		nextY := events[idx].y
		if nextY-y <= eps {
			continue
		}

		// cur is the number of crossings for any scanline with offset in [y, nextY)
		if cur > bestK {
			bestK = cur
			bestC = (y + nextY) / 2.0 // pick a representative offset inside the max interval
		}
	}

	return bestK, bestC, true
}

// Handy for visualization: return two far-apart points on the maximizing line
// so you can draw it as a segment. The line is dot(w,x)=c, direction is u.
func LineSegmentForPlot(poly []Point, u, w Point, c float64) (p0, p1 Point) {
	// pick pBase on the line: pBase = w*c (since w is unit)
	pBase := mul(w, c)

	// choose a length big enough to cross the polygon bbox
	minX, maxX := poly[0].X, poly[0].X
	minY, maxY := poly[0].Y, poly[0].Y
	for _, p := range poly[1:] {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	diag := math.Hypot(maxX-minX, maxY-minY)
	L := 10 * diag
	if L == 0 {
		L = 1
	}

	p0 = add(pBase, mul(u, -L))
	p1 = add(pBase, mul(u, +L))
	return
}

// /
func SimplePolygonGenerator(arraySize int, startingRange int, endingRange int) func(rng propcheck.SimpleRNG) ([]Point, propcheck.SimpleRNG) {
	xs := propcheck.ChooseArray(arraySize, arraySize, propcheck.ChooseInt(startingRange, endingRange))
	ys := propcheck.ChooseArray(arraySize, arraySize, propcheck.ChooseInt(startingRange, endingRange))

	ps := propcheck.Map2(xs, ys, func(xs, ys []int) []Point {
		r := make([]Point, len(xs))
		for i, x := range xs {
			p := Point{float64(x), float64(ys[i])}
			r[i] = p
		}
		lt := func(l, r Point) bool {
			if l.X < r.X {
				return true
			}
			if l.X == r.X && l.Y < r.Y {
				return true
			}
			return false

		}
		eq := func(l, r Point) bool {
			if l.X == r.X && l.Y == r.Y {
				return true
			}
			return false

		}
		s := sets.ToSet(r, lt, eq)
		r, err := BuildSimplePolygon(s) //This function written by chatGPT
		fmt.Println(err)
		return r
	})
	return ps
}

func TestMostWallPenetrationsThroughSimplePolygon(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	res := SimplePolygonGenerator(500, 1, 500)
	maxPenetrations := func(xs []Point) []Point {
		fmt.Printf("Generated array of length:%v [%v]\n", len(xs), xs) // This is what I write now.
		bestU, bestW, bestC, bestK := MaxPenetrationLine(xs)
		p0, p1 := LineSegmentForPlot(xs, bestU, bestW, bestC)
		fmt.Printf("Best max penetrations bestu %v\n", bestU)
		fmt.Printf("Best max penetrations bestW %v\n", bestW)
		fmt.Printf("Best max penetrations bestC %v\n", bestC)
		fmt.Printf("Best max penetrations bestk %v\n", bestK)
		fmt.Printf("Line Segment for plot po:%v p1:%v \n", p0, p1)
		return xs
	}
	verifyMaxPenetrations := func(actual []Point) (bool, error) {
		expected := make([]Point, len(actual))
		copy(expected, actual)
		return true, nil
	}
	test := propcheck.ForAll(res, "Find a line going through the polygon that crosses the most number of edges.", maxPenetrations, verifyMaxPenetrations)
	propcheck.ExpectSuccess[[]Point](t, test.Run(propcheck.RunParms{100, rng}))
}

const eps = 1e-12

var (
	ErrTooFewPoints = errors.New("need at least 3 unique points")
	ErrNoConverge   = errors.New("did not converge; input may be degenerate (duplicates/collinear)")
)

// BuildSimplePolygon returns vertices in cycle order (does not repeat the first point at the end).
// Uses every input point (after near-duplicate removal) exactly once.
func BuildSimplePolygon(pts []Point) ([]Point, error) {
	pts = uniquePoints(pts)
	n := len(pts)
	if n < 3 {
		return nil, ErrTooFewPoints
	}

	// 1) Initial tour: polar angle sort about centroid (deterministic, decent start).
	c := centroid(pts)
	sort.Slice(pts, func(i, j int) bool {
		ai := math.Atan2(pts[i].Y-c.Y, pts[i].X-c.X)
		aj := math.Atan2(pts[j].Y-c.Y, pts[j].X-c.X)
		if almostEqual(ai, aj) {
			// tie-breaker by distance to centroid
			return dist2(pts[i], c) < dist2(pts[j], c)
		}
		return ai < aj
	})

	// 2) 2-opt until no intersections.
	// Naive intersection scan is O(n^2) per improvement; fine for a few thousand points.
	maxIters := n * n * 20
	for iter := 0; iter < maxIters; iter++ {
		i, j, ok := findAnyCrossing(pts)
		if !ok {
			return pts, nil // simple polygon achieved
		}
		// Reverse (i+1 ... j) to remove crossing between edges (i,i+1) and (j,j+1).
		reverse(pts[i+1 : j+1])
	}

	return nil, ErrNoConverge
}

// --- crossing detection ---

func findAnyCrossing(poly []Point) (int, int, bool) {
	n := len(poly)
	for i := 0; i < n; i++ {
		a1 := poly[i]
		a2 := poly[(i+1)%n]
		for j := i + 1; j < n; j++ {
			// Skip same/adjacent edges (share endpoints)
			if edgesAdjacent(i, j, n) {
				continue
			}
			b1 := poly[j]
			b2 := poly[(j+1)%n]
			if segmentsIntersect(a1, a2, b1, b2) {
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

func edgesAdjacent(i, j, n int) bool {
	if i == j {
		return true
	}
	if (i+1)%n == j || (j+1)%n == i {
		return true
	}
	// wrap adjacency: edge (n-1->0) adjacent to (0->1)
	if (i == 0 && j == n-1) || (j == 0 && i == n-1) {
		return true
	}
	return false
}

// segmentsIntersect returns true for proper crossings and for collinear overlap/touching.
// Treating touch/overlap as intersection helps avoid degenerate “almost simple” outputs.
func segmentsIntersect(a, b, c, d Point) bool {
	if !bboxOverlap(a, b, c, d) {
		return false
	}

	o1 := orient(a, b, c)
	o2 := orient(a, b, d)
	o3 := orient(c, d, a)
	o4 := orient(c, d, b)

	// Proper intersection: strict straddle
	if sign(o1)*sign(o2) < 0 && sign(o3)*sign(o4) < 0 {
		return true
	}

	// Collinear/touching cases
	if almostZero(o1) && onSegment(a, b, c) {
		return true
	}
	if almostZero(o2) && onSegment(a, b, d) {
		return true
	}
	if almostZero(o3) && onSegment(c, d, a) {
		return true
	}
	if almostZero(o4) && onSegment(c, d, b) {
		return true
	}

	return false
}

func orient(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

func bboxOverlap(a, b, c, d Point) bool {
	minAx, maxAx := math.Min(a.X, b.X), math.Max(a.X, b.X)
	minAy, maxAy := math.Min(a.Y, b.Y), math.Max(a.Y, b.Y)
	minCx, maxCx := math.Min(c.X, d.X), math.Max(c.X, d.X)
	minCy, maxCy := math.Min(c.Y, d.Y), math.Max(c.Y, d.Y)

	return minAx <= maxCx+eps && maxAx+eps >= minCx &&
		minAy <= maxCy+eps && maxAy+eps >= minCy
}

func onSegment(a, b, p Point) bool {
	return p.X >= math.Min(a.X, b.X)-eps && p.X <= math.Max(a.X, b.X)+eps &&
		p.Y >= math.Min(a.Y, b.Y)-eps && p.Y <= math.Max(a.Y, b.Y)+eps
}

func sign(x float64) int {
	if x > eps {
		return 1
	}
	if x < -eps {
		return -1
	}
	return 0
}

func almostZero(x float64) bool     { return math.Abs(x) <= eps }
func almostEqual(a, b float64) bool { return math.Abs(a-b) <= eps }

// --- helpers ---

func reverse(s []Point) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func centroid(pts []Point) Point {
	var sx, sy float64
	for _, p := range pts {
		sx += p.X
		sy += p.Y
	}
	n := float64(len(pts))
	return Point{X: sx / n, Y: sy / n}
}

func dist2(a, b Point) float64 {
	dx, dy := a.X-b.X, a.Y-b.Y
	return dx*dx + dy*dy
}

// uniquePoints removes near-duplicates by snapping to an eps grid.
// If you truly need to keep extremely close points distinct, lower eps.
func uniquePoints(pts []Point) []Point {
	type key struct{ x, y int64 }
	m := make(map[key]Point, len(pts))
	for _, p := range pts {
		k := key{
			x: int64(math.Round(p.X / eps)),
			y: int64(math.Round(p.Y / eps)),
		}
		if _, ok := m[k]; !ok {
			m[k] = p
		}
	}
	out := make([]Point, 0, len(m))
	for _, p := range m {
		out = append(out, p)
	}
	return out
}
