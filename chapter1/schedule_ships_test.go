package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"math/rand"
	"testing"
	"time"
)

func shuffle(toBeShuffled []int) []int {
	rr := make([]int, len(toBeShuffled))
	copy(rr, toBeShuffled)
	var rr2 []int
	for _, x := range rr {
		rr2 = append(rr2, x)
		rr2 = append(rr2, 0)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rr2), func(i, j int) {
		rr2[i], rr2[j] = rr2[j], rr2[i]
	})
	return rr2
}

func TestShipSchedule(t *testing.T) {
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

	g0 := propcheck.ChooseInt(1, 3000)
	g1 := sets.ChooseSet(0, 50, g0, lt, eq)
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	prop := propcheck.ForAll(g1,
		"Validate Ship scheduling algorithm  \n",
		func(xs []int) []*Ship {
			var r []*Ship
			for _, x := range xs {
				s := Ship{
					Id:               x,
					ProposedSchedule: shuffle(xs),
					ActualSchedule:   nil,
				}
				r = append(r, &s)
			}
			return r
		},
		func(ships []*Ship) (bool, error) {
			var errors error
			start := time.Now()
			r := schedule(ships)
			fmt.Printf("Scheduling an array of %v ships took %v\n", len(ships), time.Since(start))
			for i, ship := range r { //Range loop over array of all ships
				currentShipSchedule := ship.ActualSchedule
				for j := i + 1; j < len(ships); j++ { //For current ship iterate over all ships later in array and truncate ship's Proposed Schedule at earliest conflict
					otherShipSchedule := ships[j].ActualSchedule
					for k, _ := range otherShipSchedule {
						if k < len(currentShipSchedule)-1 && currentShipSchedule[k] > 0 { //Current ship not at sea and is not truncated before otherShip
							if currentShipSchedule[k] == otherShipSchedule[k] { //Not at same port on same day
								errors = multierror.Append(errors, fmt.Errorf("Ship:%v scheduled port:%v conflicted with ship%v on day:%v", ship.Id, currentShipSchedule[k], ships[j].Id, k))

							}
						}
					}
				}
			}
			//		fmt.Println(r)
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{200, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
