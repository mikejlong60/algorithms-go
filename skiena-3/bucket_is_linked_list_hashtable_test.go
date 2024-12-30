package skiena_3

import (
	"bitbucket.org/pcastools/hash"
	"fmt"
	"github.com/greymatter-io/golangz/option"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestGetAndSet(t *testing.T) {
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	f := func(a []int, b []string, c int) propcheck.Pair[[]KeyValuePair[int32, string], int32] {
		var r = make([]KeyValuePair[int32, string], len(a))
		for c, d := range a {
			r[c] = KeyValuePair[int32, string]{int32(d), b[c]}
		}
		s := propcheck.Pair[[]KeyValuePair[int32, string], int32]{r, int32(c)}
		return s
	}

	ge := propcheck.ChooseArray(50000, 50000, propcheck.ChooseInt(-100000, 100000))
	gf := propcheck.ChooseArray(50000, 50000, propcheck.String(40))
	gh := propcheck.ChooseInt(1, 500000)
	gg := propcheck.Map3(ge, gf, gh, f)

	prop := propcheck.ForAll(gg,
		"Test get and set when you have few hash collisions  \n",
		func(xs propcheck.Pair[[]KeyValuePair[int32, string], int32]) propcheck.Pair[[]KeyValuePair[int32, string], int32] {
			return xs
		},
		func(xss propcheck.Pair[[]KeyValuePair[int32, string], int32]) (bool, error) {
			fFNV32a := func(x int32) uint32 {
				return hash.Int32(x)
			}
			eq := func(k1, k2 KeyValuePair[int32, string]) bool {
				if k1.key == k2.key {
					return true
				} else {
					return false
				}
			}
			m := New2[int32, string](eq, fFNV32a, xss.B)
			start := time.Now()
			for _, b := range xss.A {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				m = Set2(m, b, p)
			}
			fmt.Printf("Inserting %v values into hashmap with %v buckets took:%v\n", len(xss.A), xss.B, time.Since(start))

			var errors error

			//The value should be in the map
			start = time.Now()
			for _, b := range xss.A {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				err := option.GetOrElse(Get2(m, b.key, p), KeyValuePair[int32, string]{b.key, fmt.Sprintf("Should have found:%v in HashMap", 188)}) //fmt.Sprintf("Should not have found:%v in HashMap", 188))
				if err.key != b.key {
					errors = multierror.Append(errors, fmt.Errorf("Should have found:%v in HashMap", b.key))
				}
			}
			fmt.Printf("Getting %v values from hashmap with %v buckets took:%v\n", len(xss.A), xss.B, time.Since(start))

			//////////  Now use Golang's hashmap to compare performance
			mg := make(map[int32]KeyValuePair[int32, string], xss.B)
			start = time.Now()
			for _, b := range xss.A {
				mg[b.key] = b
			}
			fmt.Printf("Inserting %v values into Golang hashmap took:%v\n", len(xss.A), time.Since(start))
			//The value should be in the map
			var ss KeyValuePair[int32, string]
			start = time.Now()
			for _, b := range xss.A {
				ss = mg[b.key]
			}
			fmt.Printf("Getting %v values from Golang hashmap %v took:%v\n", len(xss.A), ss, time.Since(start))

			/////////////

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[[]KeyValuePair[int32, string], int32]](t, result)
}

func TestDelete(t *testing.T) {
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	f := func(a []int, b []string, c int) propcheck.Pair[[]KeyValuePair[int32, string], int32] {
		var r = make([]KeyValuePair[int32, string], len(a))
		for c, d := range a {
			r[c] = KeyValuePair[int32, string]{int32(d), b[c]}
		}
		s := propcheck.Pair[[]KeyValuePair[int32, string], int32]{r, int32(c)}
		return s
	}

	ge := propcheck.ChooseArray(50000, 50000, propcheck.ChooseInt(-100000, 100000))
	gf := propcheck.ChooseArray(50000, 50000, propcheck.String(40))
	gh := propcheck.ChooseInt(1, 500000)
	gg := propcheck.Map3(ge, gf, gh, f)

	prop := propcheck.ForAll(gg,
		"Test delete with random number of buckets that may cause hash collisions  \n",
		func(xs propcheck.Pair[[]KeyValuePair[int32, string], int32]) propcheck.Pair[[]KeyValuePair[int32, string], int32] {
			return xs
		},
		func(xss propcheck.Pair[[]KeyValuePair[int32, string], int32]) (bool, error) {
			fFNV32a := func(x int32) uint32 {
				return hash.Int32(x)
			}
			eq := func(k1, k2 KeyValuePair[int32, string]) bool {
				if k1.key == k2.key {
					return true
				} else {
					return false
				}
			}
			m := New2[int32, string](eq, fFNV32a, xss.B)
			start := time.Now()
			for _, b := range xss.A {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				m = Set2(m, b, p)
			}
			fmt.Printf("Inserting %v values into hashmap with %v buckets took:%v\n", len(xss.A), xss.B, time.Since(start))

			var errors error

			//The value should be in the map
			start = time.Now()
			for _, b := range xss.A {
				p := func(s KeyValuePair[int32, string]) bool {
					if s.key == b.key { //Close around the variable in the loop
						return true
					} else {
						return false
					}
				}
				newHashmap := Delete2(m, b.key, p)
				err := option.GetOrElse(Get2(newHashmap, b.key, p), KeyValuePair[int32, string]{-10000000, fmt.Sprintf("Should not have found key:%v in HashMap", b.key)})
				if err.key == b.key {
					errors = multierror.Append(errors, fmt.Errorf("Should not have found:%v in HashMap", b.key))
				}
			}
			fmt.Printf("Deleting %v values from hashmap that had %v buckets took:%v\n", len(xss.A), xss.B, time.Since(start))

			//////////  Now use Golang's hashmap to compare performance
			mg := make(map[int32]KeyValuePair[int32, string])
			start = time.Now()
			for _, b := range xss.A {
				mg[b.key] = b
			}
			fmt.Printf("Inserting %v values into Golang hashmap took:%v\n", len(xss.A), time.Since(start))
			//The value should be in the map
			start = time.Now()
			for _, b := range xss.A {
				delete(mg, b.key)
			}
			fmt.Printf("Deleting %v values from Golang hashmap  took:%v\n", len(xss.A), time.Since(start))

			/////////////

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[[]KeyValuePair[int32, string], int32]](t, result)
}
