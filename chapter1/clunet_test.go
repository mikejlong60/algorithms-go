package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestClunetSwitch(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}

	g0 := propcheck.ChooseInt(1, 300)
	g1 := sets.ChooseSet(0, 100, g0, lt, eq)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g1,
		"Validate Clunet switch algorithm  \n",
		func(xs []int) propcheck.Pair[[]*InputWire, []*OutputWire] {
			var outputWires []*OutputWire
			for _, y := range xs {
				ow := OutputWire{
					Id:             y,
					InputJunctions: make([]*InputWire, 1, 1),
				}
				outputWires = append(outputWires, &ow)
			}
			var inputWires []*InputWire
			for _, x := range xs {
				s := InputWire{
					Id: x,
				}
				inputWires = append(inputWires, &s)
			}
			return propcheck.Pair[[]*InputWire, []*OutputWire]{inputWires, outputWires}
		},
		func(wires propcheck.Pair[[]*InputWire, []*OutputWire]) (bool, error) {

			eq := func(l, r *InputWire) bool {
				if l.Id == r.Id {
					return true
				} else {
					return false
				}
			}

			lt := func(l, r *InputWire) bool {
				if l.Id < r.Id {
					return true
				} else {
					return false
				}
			}
			var errors error
			start := time.Now()
			r := MakeSwitches(wires)
			fmt.Printf("Scheduling an array of %v wires took %v\n", len(wires.A), time.Since(start))
			var liw []*InputWire

			//The switching operation ensures that
			//  1. Every output wire contains every input wire
			//  2. Every output wire switches on a different input wire
			//  3. No input wires passes in front of any other input wire on an output wire on the way to its switch. This is accomplished by switching on the first junction of every output wire.
			for _, ow := range r {
				liw = append(liw, ow.InputJunctions[0])
			}
			liwAsSet := sets.ToSet(liw, lt, eq)
			if len(liwAsSet) != len(r) {
				errors = multierror.Append(errors, fmt.Errorf("Expected the length:%v of the set of last input junctions for all output wires to equal the size of the set resulting from the switching operation:%v", len(liwAsSet), len(r)))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
