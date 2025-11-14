package chapter1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
)

func shuffleAny[A any](toBeShuffled []*A) []*A {
	rr := make([]*A, len(toBeShuffled))
	copy(rr, toBeShuffled)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rr), func(i, j int) {
		rr[i], rr[j] = rr[j], rr[i]
	})
	return rr
}

func TestStableMatchingPropertiesTest(t *testing.T) {
	var allTheMen []*Man
	var allTheWomen []*Woman
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	g0 := propcheck.ChooseInt(0, 1000)
	fa := func(a int) func(propcheck.SimpleRNG) (*linked_list.LinkedList[*Man], propcheck.SimpleRNG) {
		allTheMen = []*Man{}
		allTheWomen = []*Woman{}
		mw := func(mIds []int) *linked_list.LinkedList[*Man] {
			if len(allTheWomen) != len(allTheMen) {
				t.Error("length of men and women arrays were not equal")
			}
			for _, s := range mIds {
				allTheMen = append(allTheMen, &Man{Id: s})
			}
			for _, s := range mIds {
				allTheWomen = append(allTheWomen, &Woman{Id: s})
			}

			//Make two arrays, one of shuffled men, one for each woman, and one of  shuffled women, one for each man.
			var freeMen *linked_list.LinkedList[*Man]
			for _, s := range allTheMen {
				freeMen = linked_list.Push(s, freeMen)
			}

			for i, _ := range allTheWomen {
				womenForMan := shuffleAny(allTheWomen)
				var allWomen *linked_list.LinkedList[*Woman]
				for _, s := range womenForMan {
					allWomen = linked_list.Push(s, allWomen)
				}
				allTheMen[i].HaveNotProposedTo = allWomen
				allTheMen[i].Preferences = linked_list.ToArray(allWomen)
			}

			for _, s := range allTheWomen {
				wpref := shuffleAny(allTheMen)
				var wprefMap = make(map[int]propcheck.Pair[int, *Man], len(wpref))
				for i, m := range wpref {
					wprefMap[m.Id] = propcheck.Pair[int, *Man]{i, m}
				}
				s.Preferences = wprefMap
			}

			return freeMen
		}
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

		//Make these sets intead of arays
		g00 := propcheck.ChooseInt(1, 3000)
		g1 := sets.ChooseSet(0, 1000, g00, lt, eq)
		return propcheck.Map(g1, mw)
	}

	g := propcheck.FlatMap(g0, fa)

	prop := propcheck.ForAll(g,
		"Make a bunch of men and women and match them up and see if all matches are stable \n",
		func(freeMen *linked_list.LinkedList[*Man]) []*Woman {
			len := linked_list.Len(freeMen)
			start := time.Now()
			r := Match(freeMen)
			fmt.Printf("Match took %v for %v couples\n", time.Since(start), len)
			return r
		},
		func(allWomen []*Woman) (bool, error) {
			var errors error

			var allHusbandIds []int
			for _, j := range allWomen {
				if j.EngagedTo == nil {
					errors = multierror.Append(errors, fmt.Errorf("Woman:%v was not married ", j.Id))
				}
				allHusbandIds = append(allHusbandIds, j.EngagedTo.Id)
			}

			var allMenIds []int
			for _, man := range allTheMen {
				allMenIds = append(allMenIds, man.Id)
			}

			mEq := func(m1, m2 int) bool {
				if m1 == m2 {
					return true
				} else {
					return false
				}
			}
			if !sets.SetEquality(allMenIds, allHusbandIds, mEq) {
				errors = multierror.Append(errors, fmt.Errorf("All men were not married"))
				fmt.Printf("These men never got married:%v\n", sets.SetMinus(allMenIds, allHusbandIds, mEq))
			}
			unstableMatchings := unstableMatchings(allWomen)
			if len(unstableMatchings) > 0 {
				errors = multierror.Append(errors, fmt.Errorf("There were unstable matchings as follow:%v", unstableMatchings))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{20, rng})
	propcheck.ExpectSuccess[*linked_list.LinkedList[*Man]](t, result)
}

func TestStableMatchingPropertiesIndifferenceTest(t *testing.T) {
	var allTheMen []*Man
	var allTheWomen []*Woman
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	g0 := propcheck.ChooseInt(0, 3000)
	fa := func(a int) func(propcheck.SimpleRNG) (*linked_list.LinkedList[*Man], propcheck.SimpleRNG) {
		allTheMen = []*Man{}
		allTheWomen = []*Woman{}
		mw := func(mIds []int) *linked_list.LinkedList[*Man] {
			if len(allTheWomen) != len(allTheMen) {
				t.Error("length of men and women arrays were not equal")
			}
			for _, s := range mIds {
				allTheMen = append(allTheMen, &Man{Id: s})
			}
			for _, s := range mIds {
				allTheWomen = append(allTheWomen, &Woman{Id: s})
			}

			//Make two arrays, one of shuffled men, one for each woman, and one of  shuffled women, one for each man.
			var freeMen *linked_list.LinkedList[*Man]
			for _, s := range allTheMen {
				freeMen = linked_list.Push(s, freeMen)
			}

			for i, _ := range allTheWomen {
				womenForMan := shuffleAny(allTheWomen)
				var allWomen *linked_list.LinkedList[*Woman]
				for _, s := range womenForMan {
					allWomen = linked_list.Push(s, allWomen)
				}
				allTheMen[i].HaveNotProposedTo = allWomen
				allTheMen[i].Preferences = linked_list.ToArray(allWomen)
			}

			for _, s := range allTheWomen {
				wpref := shuffleAny(allTheMen)
				var wprefMap = make(map[int]propcheck.Pair[int, *Man], len(wpref)/2)
				for i, m := range wpref {
					wprefMap[m.Id] = propcheck.Pair[int, *Man]{i, m}
					if i > len(wpref)/2 { //Make every woman indifferent to half the men by excluding them from her preference list
						break
					}
				}
				s.Preferences = wprefMap
			}

			return freeMen
		}
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

		//Make these sets intead of arays
		g00 := propcheck.ChooseInt(1, 3000)
		g1 := sets.ChooseSet(0, 1000, g00, lt, eq)
		return propcheck.Map(g1, mw)
	}

	g := propcheck.FlatMap(g0, fa)

	prop := propcheck.ForAll(g,
		"Make a bunch of men and women and match them up and see if all matches are stable \n",
		func(freeMen *linked_list.LinkedList[*Man]) []*Woman {
			ln := linked_list.Len(freeMen)
			start := time.Now()
			r := Match(freeMen)
			fmt.Printf("Match took %v for %v couples\n", time.Since(start), ln)
			return r
		},
		func(allWomen []*Woman) (bool, error) {
			var errors error

			var allHusbandIds []int
			for _, j := range allWomen {
				if j.EngagedTo == nil {
					errors = multierror.Append(errors, fmt.Errorf("Woman:%v was not married ", j.Id))
				}
				allHusbandIds = append(allHusbandIds, j.EngagedTo.Id)
			}

			var allMenIds []int
			for _, man := range allTheMen {
				allMenIds = append(allMenIds, man.Id)
			}

			mEq := func(m1, m2 int) bool {
				if m1 == m2 {
					return true
				} else {
					return false
				}
			}
			if !sets.SetEquality(allMenIds, allHusbandIds, mEq) {
				errors = multierror.Append(errors, fmt.Errorf("All men were not married"))
				fmt.Printf("These men never got married:%v\n", sets.SetMinus(allMenIds, allHusbandIds, mEq))
			}
			unstableMatchings := unstableMatchings(allWomen)
			if len(unstableMatchings) > 0 {
				errors = multierror.Append(errors, fmt.Errorf("There were unstable matchings as follow:%v", unstableMatchings))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{20, rng})
	propcheck.ExpectSuccess[*linked_list.LinkedList[*Man]](t, result)
}

//Algorithm for determining unstable matchings
//   Make a new list r of strings that you will return as the list of Unstable Matchings
//   For1 each woman w from all women
//        grab w's husband as m
//        make a new list ipw of potential instabilities
//        For2 each man m2 in woman w's preferences
//            if m2 ranks w above his current woman w2?
//                add m2 to ipw as a candidate for instability for woman w
//            end if
//        End For2
//        For3 each ipw as m3
//            if m3 ranks above w's current husband
//                 log that notation that w prefers m3 and m3 prefers w. This is in comparison to m.  Log all three of these values
//            end if
//        End For3
//    End For1

//If len(r) > 1 you have unstable matching and you should print list to console

func unstableMatchings(allWomen []*Woman) []string {
	var unstableMatchings []string
	var womanRanking = func(w *Woman, currentWomanPreferences []*Woman) int {
		for k, w2 := range currentWomanPreferences {
			if w.Id == w2.Id {
				return k
			}
		}
		return len(currentWomanPreferences) - 1 //Return lowest possible ranking if man has no preference for his current woman
	}
	for _, w := range allWomen { //for1
		m := w.EngagedTo
		var ipw []*Man
		var wPrefs []int
		for _, k := range w.Preferences {
			wPrefs = append(wPrefs, k.B.Id)
		}
		for _, m2 := range w.Preferences { //for2
			//Get man for m2 id
			//Then determine for each man if new woman w ranks above his current woman in his preferences list versus the woman to whom he is currently engaged and add that
			//man to the list of instability candidates for evaluation in the next for loop.
			mansRankingOfCurrentWoman := womanRanking(w, m2.B.Preferences)
			mansRankingOfCurrentSpouse := womanRanking(m2.B.EngagedTo, m2.B.Preferences)
			//fmt.Printf("\n\nMan %v ranking of current woman %v is:%v --- ", m2.B.Id, w.Id, mansRankingOfCurrentWoman)
			//fmt.Printf("versus his ranking of current spouse:%v", mansRankingOfCurrentSpouse)
			if mansRankingOfCurrentWoman < mansRankingOfCurrentSpouse { //lower is a higher preference
				//fmt.Printf(" adding him as an instability candidate\n")
				ipw = append(ipw, m2.B)
			}
		}
		//end for2
		for _, m3 := range ipw { //for3
			if w.Preferences[m3.Id].A < w.Preferences[m.Id].A { //lower on the preference list is better
				//fmt.Printf("\nWoman:%v prefers Man:%v over her current husband:%v and this is an instablity\n", w.Id, m3.Id, m.Id)
				unstableMatchings = append(unstableMatchings, fmt.Sprintf("Woman:%v prefers Man:%v over her current husband:%v and this is an instablity", w.Id, m3.Id, m.Id))
			}
		} //end for3
	} //end for1
	return unstableMatchings
}
