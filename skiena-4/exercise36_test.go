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

// Generates an n array of A[1..n] that has numbers between 1 and n(squared), but which contains at most
// log log n different numbers.  The array is a size between start and stopExclusive
func ArrayOfNWithBase2LogLogNDifferentValues(start int, stopExclusive int) func(propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
	g1 := propcheck.ChooseInt(start, stopExclusive)
	//make it a float
	g2 := propcheck.Map(g1, func(x int) float64 {
		return float64(x)
	})

	// Take that number and get its log2 log2(a) and generate an array of x numbers
	// with only a different randomly generated numbers.
	g3 := func(x float64) func(propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
		return func(rng propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {

			r := int(math.Log2(math.Log2(x)))
			fmt.Println(x)
			fmt.Println(r)
			var lr = rng //The ever-changing random number generator inside the loop below.
			var res = make([]int, 0)
			var a int
			var j int
			var start int
			for i := 0; i < r; i++ {
				a, lr = propcheck.ChooseInt(0, int(x)/r)(lr)
				for j = 0; j < a; j++ {
					res = append(res, a)
					start++
				}
			}
			return res, lr
		}
	}
	return propcheck.FlatMap(g2, g3)
}

func TestGeneratorForArrayOfNWithBase2LogLogNDifferentValues(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	res := ArrayOfNWithBase2LogLogNDifferentValues(10000, 1000000)
	actual, _ := res(rng)
	fmt.Println(actual)
	//if len(actual) != 13 {
	//	t.Errorf("Map should have incremented the unit value by 1 \n")
	//}
	//if rng != rng2 {
	//	t.Error("Should have gotten the same SimpleRNG because you just flatmapped over A Id generator")
	//}
}
