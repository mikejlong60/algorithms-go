package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"math"
	"testing"
	"time"
)

type point struct {
	x float64
	y float64
}

// Finds the closest point (p1) to p and return that point and its index in the points input array
func findClosestPoint(p point, points []point) (point, int) {
	distance := func(p1, p2 point) float64 {
		return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x)) + ((p2.y - p1.y) * (p2.y - p1.y))
	}
	var closestDistance = distance(p, points[0])
	var currentDistance float64
	var closestPointIdx int = 0
	for l, m := range points {
		currentDistance = distance(p, m)
		if currentDistance < closestDistance {
			closestDistance = currentDistance
			closestPointIdx = l
		}
	}
	closestPointToP := points[closestPointIdx]
	return closestPointToP, closestPointIdx
}

// points is starting array, closestPoints is result array.
func nearestNeighbor(p point, points, closestPoints []point) []point {
	if len(points) == 0 {
		return closestPoints
	} else {
		closestPointToP, closestPointIdx := findClosestPoint(p, points)
		points = append(points[:closestPointIdx], points[closestPointIdx+1:]...)
		closestPoints = append(closestPoints, closestPointToP)
		return nearestNeighbor(closestPointToP, points, closestPoints)
	}
}

func TestNearestNeighbor(t *testing.T) {
	g0 := propcheck.ArrayOfN(300, propcheck.ChooseInt(-300, 301))
	g1 := propcheck.ArrayOfN(300, propcheck.ChooseInt(-300, 301))
	g2 := propcheck.Map2(g0, g1, func(xs, ys []int) []point {
		r := []point{}
		for i := range xs {
			p := point{float64(xs[i]), float64(ys[i])}
			r = append(r, p)
		}
		return r
	})
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g2,
		"Validate traveling salesman algo  \n",
		func(xs []point) []point {

			return xs
		},
		func(xs []point) (bool, error) {
			var errors error
			expectedLen := len(xs)
			//expected := xs[0] * xs[1]
			//actual := xs[3]
			//if actual != expected {
			//	t.Errorf("Actual:%v Expected:%v", actual, expected)
			//}
			fmt.Printf("Origin:%v\n", xs)
			actual := nearestNeighbor(xs[0], xs[1:], []point{xs[0]}) //Starting point is first element in array.
			actualLen := len(actual)
			fmt.Printf("Actual:%v\n", actual)
			if actualLen != expectedLen {
				errors = fmt.Errorf("Actual length:%v Expected length:%v", actualLen, expectedLen)
			}

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]point](t, result)
}
