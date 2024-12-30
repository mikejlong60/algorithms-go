package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func frequencyLt(l, r *Frequency) bool {
	if l.probability < r.probability {
		return true
	} else {
		return false
	}
}

func frequencyEq(l, r *Frequency) bool {
	if l.probability == r.probability {
		return true
	} else {
		return false
	}
}

func TestHuffmanHeapFromBook(t *testing.T) {
	a := Frequency{
		.32, "a", nil, nil,
	}
	b := Frequency{
		.25, "b", nil, nil,
	}
	c := Frequency{
		.20, "c", nil, nil,
	}
	d := Frequency{
		.18, "d", nil, nil,
	}
	e := Frequency{
		.05, "e", nil, nil,
	}
	f := []*Frequency{&a, &b, &c, &d, &e}
	insertIntoHeap := func(xss []*Frequency) []*Frequency {
		var r = StartHeapF(5)
		for _, x := range xss {
			r = HeapInsertF(r, x, frequencyLt)
		}
		return r
	}
	var actual = insertIntoHeap(f)
	var expected = []*Frequency{&e, &d, &b, &a, &c}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \nexpected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&d, &c, &b, &a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&c, &b, &a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&b, &a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{&a}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
	actual, _ = HeapDeleteF(actual, 0, frequencyLt)
	expected = []*Frequency{}
	if !arrays.ArrayEquality(actual, expected, frequencyEq) {
		t.Errorf("Actual:%v, \n              expected:%v", actual, expected)
	}
}

func TestHuffmanHeapFromBook2(t *testing.T) {
	a := Frequency{
		.32, "a", nil, nil,
	}
	b := Frequency{
		.25, "b", nil, nil,
	}
	c := Frequency{
		.20, "c", nil, nil,
	}
	d := Frequency{
		.18, "d", nil, nil,
	}
	e := Frequency{
		.05, "e", nil, nil,
	}
	f := []*Frequency{&a, &b, &c, &d, &e}
	insertIntoHeap := func(xss []*Frequency) []*Frequency {
		var r = StartHeapF(5)
		for _, x := range xss {
			r = HeapInsertF(r, x, frequencyLt)
		}
		return r
	}
	var freq = insertIntoHeap(f)

	freq = Huffman(freq, frequencyLt)
	if len(freq) != 1 {
		t.Errorf("Expected freq to be len 1 but was:%v", len(freq))
	}

	expected := Frequency{probability: 1, letter: "((c:(e:d)):(b:a))"}
	if !(freq[0].letter == expected.letter && freq[0].probability == 1) {
		t.Errorf("Expected freq to be a tree with the combined letters ((c:(e:d)):(b:a)) but was:%v", freq)
	}

}
