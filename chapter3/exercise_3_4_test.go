package chapter3

import (
	"fmt"
	"testing"
)

type Butterfly struct {
	Id      string
	Species int //string
}

type ButterflyPair struct {
	Type        int // the sum of the two pairs. Odd numbers are supposed different species. Even numbers are supposed same species.  Ambiguity is not part of the al;gorithm
	SameSpecies bool
}

// O value is not greater than m + n where m is the number of non-ambiguous judgements and n is the number of butterflies
// I go through the array of judgements one time and then go through it again no more than one time(usually less) until I find
// inconsistency in the two groups.
func Consistent(xs []ButterflyPair) bool {

	consistent := func(x map[bool]([]bool)) bool {
		diff := x[false]
		same := x[true]
		differentSpeciesConsistent := func(differentSpecies []bool) bool {
			var prev = false
			for _, curr := range differentSpecies {
				if prev != curr {
					return false
				}
				prev = curr
			}
			return true
		}
		sameSpeciesConsistent := func(sameSpecies []bool) bool {
			var prev = true
			for _, curr := range sameSpecies {
				if prev != curr {
					return false
				}
				prev = curr
			}
			return true
		}
		return sameSpeciesConsistent(same) && differentSpeciesConsistent(diff)
	}

	var d = make(map[bool]([]bool))
	var sameSpecies []bool
	var differentSpecies []bool
	for _, i := range xs {
		if i.Type%2 == 0 {
			sameSpecies = append(sameSpecies, i.SameSpecies)
		} else {
			differentSpecies = append(differentSpecies, i.SameSpecies)
		}
	}
	d[true] = sameSpecies
	d[false] = differentSpecies

	return consistent(d)
}

func TestFindingsConsistent(t *testing.T) {
	a0 := Butterfly{
		Id:      "a0",
		Species: 1, //"Monarch",
	}
	a1 := Butterfly{
		Id:      "a1",
		Species: 1, //"Monarch",
	}
	a2 := Butterfly{
		Id:      "a2",
		Species: 1, //"Monarch",
	}
	a3 := Butterfly{
		Id:      "a3",
		Species: 1, //"Monarch",
	}
	a4 := Butterfly{
		Id:      "a4",
		Species: 1, //"Monarch",
	}
	a5 := Butterfly{
		Id:      "a5",
		Species: 1, //"Monarch",
	}
	//b0 := Butterfly{
	//	Id:      "b0",
	//	Species: 2, //"Viceroy",
	//}
	b1 := Butterfly{
		Id:      "b1",
		Species: 2, //"Viceroy",
	}
	b2 := Butterfly{
		Id:      "b2",
		Species: 2, //"Viceroy",
	}
	b3 := Butterfly{
		Id:      "b3",
		Species: 2, //"Viceroy",
	}
	b4 := Butterfly{
		Id:      "b4",
		Species: 2, //"Viceroy",
	}
	b5 := Butterfly{
		Id:      "b5",
		Species: 2, //"Viceroy",
	}

	p1 := ButterflyPair{
		Type:        a0.Species + a1.Species,
		SameSpecies: true,
	}
	p2 := ButterflyPair{
		Type:        a0.Species + a2.Species,
		SameSpecies: true,
	}
	p3 := ButterflyPair{
		Type:        a0.Species + a3.Species,
		SameSpecies: true,
	}
	p4 := ButterflyPair{
		Type:        a0.Species + a4.Species,
		SameSpecies: true,
	}
	p5 := ButterflyPair{
		Type:        a0.Species + a5.Species,
		SameSpecies: true,
	}
	p6 := ButterflyPair{
		Type:        a0.Species + b1.Species,
		SameSpecies: false,
	}
	p7 := ButterflyPair{
		Type:        a0.Species + b2.Species,
		SameSpecies: false,
	}
	p8 := ButterflyPair{
		Type:        a0.Species + b3.Species,
		SameSpecies: false,
	}
	p9 := ButterflyPair{
		Type:        a0.Species + b4.Species,
		SameSpecies: false,
	}
	p10 := ButterflyPair{
		Type:        a0.Species + b5.Species,
		SameSpecies: false,
	}

	xs := []ButterflyPair{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10}
	actual := Consistent(xs)
	if !actual {
		t.Errorf("Species determination was consistent but algorithm said it was not consistent")
	}
	fmt.Println(actual)
}

func TestFindingsNotConsistent(t *testing.T) {
	a0 := Butterfly{
		Id:      "a0",
		Species: 1, //"Monarch",
	}
	a1 := Butterfly{
		Id:      "a1",
		Species: 1, //"Monarch",
	}
	a2 := Butterfly{
		Id:      "a2",
		Species: 1, //"Monarch",
	}
	b1 := Butterfly{
		Id:      "b0",
		Species: 2, //"Viceroy",
	}
	b2 := Butterfly{
		Id:      "b2",
		Species: 2, //"Viceroy",
	}
	b3 := Butterfly{
		Id:      "b3",
		Species: 2, //"Viceroy",
	}

	p1 := ButterflyPair{
		Type:        a0.Species + a1.Species,
		SameSpecies: true,
	}
	p2 := ButterflyPair{
		Type:        a0.Species + a2.Species,
		SameSpecies: true,
	}
	p6 := ButterflyPair{
		Type:        a0.Species + b1.Species,
		SameSpecies: false,
	}
	p7 := ButterflyPair{
		Type:        a0.Species + b2.Species,
		SameSpecies: true, //this is an inconsistency
	}
	p8 := ButterflyPair{
		Type:        a0.Species + b3.Species,
		SameSpecies: false,
	}

	xs := []ButterflyPair{p1, p2, p6, p7, p8}
	actual := Consistent(xs)
	if actual {
		t.Errorf("Species determination was not consistent but algorithm said it was consistent")
	}
	fmt.Println(actual)
}
