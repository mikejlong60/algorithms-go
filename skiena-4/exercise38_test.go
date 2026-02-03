package skiena_4

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

type Point struct {
	X, Y float64
}

// ConvexHull returns the convex hull of pts as a polygon in CCW order.
// If pts has < 3 unique points, it returns the unique points (no polygon possible).
//
// By default this version *excludes* collinear points on edges (keeps only corners).
// To *include* collinear boundary points, change the cross-check from <= 0 to < 0.
func ConvexHull(pts []Point) []Point {
	pts = uniquePoints(pts)
	n := len(pts)
	if n <= 2 {
		return pts
	}

	sort.Slice(pts, func(i, j int) bool {
		if pts[i].X == pts[j].X {
			return pts[i].Y < pts[j].Y
		}
		return pts[i].X < pts[j].X
	})

	// Build lower hull
	lower := make([]Point, 0, n)
	for _, p := range pts {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	// Build upper hull
	upper := make([]Point, 0, n)
	for i := n - 1; i >= 0; i-- {
		p := pts[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	// Concatenate lower + upper minus duplicate endpoints
	// lower: left->right, upper: right->left
	hull := append(lower[:len(lower)-1], upper[:len(upper)-1]...)

	// If all points were collinear, hull can end up with 2 points.
	return hull
}

// cross returns the z-component of (b-a) x (c-a).
// > 0 => counter-clockwise turn
// < 0 => clockwise turn
// = 0 => collinear
func cross(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

// uniquePoints removes duplicates with a tolerance (epsilon) to avoid float noise issues.
func uniquePoints(pts []Point) []Point {
	const eps = 1e-12
	type key struct{ x, y int64 }

	m := make(map[key]Point, len(pts))
	for _, p := range pts {
		k := key{
			x: int64(math.Round(p.X / eps)),
			y: int64(math.Round(p.Y / eps)),
		}
		// Keep the first representative (good enough for hull)
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

func TestConvexHullFromChatGPT(t *testing.T) {
	xs := []Point{{1.0, 1.0}, {2.0, 5.0}, {3.0, 2.0}, {5.0, 4.0}, {5, 5}, {3, 4}, {3, 4}}
	r := ConvexHull(xs)
	fmt.Println(r)
}
