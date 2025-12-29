package skiena_4

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/propcheck"
)

// Given an array of A[1..n] that has numbers between 1 and n(squared), but which contains at most
// log log n different numbers, show that you can sort it in less than O(n log n).

func TestGeneratorForArrayOfNWithBase2LogLogNDifferentValues(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	g1 := propcheck.ChooseInt(10000, 1000000)
	//make it a float
	g2 := propcheck.Map(g1, func(x int) float64 {
		return float64(x)
	})

	///return func(rng SimpleRNG) (string, SimpleRNG)

	g3 := func(x float64, y propcheck.SimpleRNG) func(propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
		return func(rng propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
			fmt.Println(x)

			r := int(math.Log2(math.Log2(x)))
			var lr = rng //The ever-changing random number generator inside the loop below.
			var res = make([]int, int(x))
			var a int
			for i := 0; i < r; i++ {
				a, lr = propcheck.ChooseInt(0, int(x)/r)(lr)
				a1 := make([]int, a/3)
				for j := 0; j < a; j++ {
					a1 = append(a1, a)
				}
				res = append(res, a1...)
			}
			return res, lr
		}
	}
	res := propcheck.FlatMap(g2, g3)
	actual, _ := res(rng)
	if len(actual) != 13 {
		t.Errorf("Map should have incremented the unit value by 1 \n")
	}
	//if rng != rng2 {
	//	t.Error("Should have gotten the same SimpleRNG because you just flatmapped over A Id generator")
	//}
}
