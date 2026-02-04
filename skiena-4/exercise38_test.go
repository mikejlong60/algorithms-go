package skiena_4

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

func SimplePolygonGenerator(arraySize int, start int, stopExclusive int) func(rng propcheck.SimpleRNG) ([]Point, propcheck.SimpleRNG) {
	xs := propcheck.ChooseArray(arraySize, arraySize, propcheck.ChooseInt(start, stopExclusive))
	ys := propcheck.ChooseArray(arraySize, arraySize, propcheck.ChooseInt(start, stopExclusive))

	ps := propcheck.Map2(xs, ys, func(xs, ys []int) []Point {
		r := make([]Point, len(xs))
		for x, i := range xs {
			p := Point{float64(x), float64(ys[i])}
			r[i] = p
		}
		lt := func(l, r Point) bool {
			if l.X > r.X {
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
		return sets.ToSet(r, lt, eq)
	})
	return ps
}

func TestSimplePolygon(t *testing.T) {
	xs := []Point{{1.0, 1.0}, {2.0, 5.0}, {3.0, 2.0}, {5.0, 4.0}, {5, 5}, {3, 4}, {4, 4}}
	r, err := BuildSimplePolygon(xs)
	fmt.Println(r)
	fmt.Println(err)
}

type Point struct {
	X, Y float64
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
