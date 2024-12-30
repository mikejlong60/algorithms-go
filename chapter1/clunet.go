package chapter1

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

type InputWire struct {
	Id int
}

type OutputWire struct {
	Id             int
	InputJunctions []*InputWire
}

func (w InputWire) String() string {
	return fmt.Sprintf("InputWire{Id:%v}", w.Id)
}

func (w OutputWire) String() string {
	var iwIds = make([]int, len(w.InputJunctions), len(w.InputJunctions))
	for i, iw := range w.InputJunctions {
		iwIds[i] = iw.Id
	}
	return fmt.Sprintf("OutputWire{Id:%v, InputJunctions:%v}", w.Id, iwIds)
}

func MakeSwitches(wires propcheck.Pair[[]*InputWire, []*OutputWire]) []*OutputWire {
	iwEq := func(l, r *InputWire) bool {
		if l == nil {
			return false
		} else if l.Id == r.Id {
			return true
		} else {
			return false
		}
	}

	if len(wires.A) != len(wires.B) {
		panic("input wires array must be same size as outputwires array")
	}
	if len(wires.A) == 0 {
		return []*OutputWire{}
	}

	//Add next input wire as first junction on every next output wire
	for i, iw := range wires.A {
		ow := wires.B[i]
		ow.InputJunctions[0] = iw
	}

	for _, ow := range wires.B {
		otherJunctions := sets.SetMinus(wires.A, ow.InputJunctions, iwEq)
		ow.InputJunctions = arrays.Append(otherJunctions, ow.InputJunctions)
	}
	return wires.B
}
