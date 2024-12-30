package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func TestKruskals(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 2,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 30,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 31,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 3,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, bd, cd, cb})
	expected := []*PrimsEdge{ac, cb, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskalsUsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 2,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 30,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 31,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 3,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, bd, cd, cb})
	expected := []*PrimsEdge{ac, cb, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals1(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		weight: 4.3,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 4.4,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 4.5,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 4.6,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ad, ab}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals1UsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		weight: 4.3,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 4.4,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 4.5,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 4.6,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ad, ab}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals2(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		weight: 1,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 4.3,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 4.4,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 4.5,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ab, ad}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals2UsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4.1,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 4.2,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		weight: 1,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 4.3,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 4.4,
	}
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 4.5,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, ad, bd, cd, cb})
	expected := []*PrimsEdge{ac, ab, ad}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals3(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 8,
	}

	bc := &PrimsEdge{
		u:      "b",
		v:      "c",
		weight: 11,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 8.1,
	}

	cf := &PrimsEdge{
		u:      "c",
		v:      "f",
		weight: 1,
	}

	ce := &PrimsEdge{
		u:      "c",
		v:      "e",
		weight: 7.1,
	}
	de := &PrimsEdge{
		u:      "d",
		v:      "e",
		weight: 2,
	}
	dh := &PrimsEdge{
		u:      "d",
		v:      "h",
		weight: 4.1,
	}
	dg := &PrimsEdge{
		u:      "d",
		v:      "g",
		weight: 7.2,
	}

	ef := &PrimsEdge{
		u:      "e",
		v:      "f",
		weight: 6,
	}
	fh := &PrimsEdge{
		u:      "f",
		v:      "h",
		weight: 2.1,
	}

	hi := &PrimsEdge{
		u:      "h",
		v:      "i",
		weight: 10,
	}
	gh := &PrimsEdge{
		u:      "g",
		v:      "h",
		weight: 14,
	}
	gi := &PrimsEdge{
		u:      "g",
		v:      "i",
		weight: 9,
	}
	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := Kruskals([]*PrimsEdge{ab, ac, bc, bd, cf, ce, de, dh, dg, ef, fh, hi, gi, gh})
	expected := []*PrimsEdge{ab, ac, cf, de, dh, dg, gi, fh}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestKruskals3UsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4,
	}
	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 8,
	}

	bc := &PrimsEdge{
		u:      "b",
		v:      "c",
		weight: 11,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 8.1,
	}

	cf := &PrimsEdge{
		u:      "c",
		v:      "f",
		weight: 1,
	}

	ce := &PrimsEdge{
		u:      "c",
		v:      "e",
		weight: 7.1,
	}
	de := &PrimsEdge{
		u:      "d",
		v:      "e",
		weight: 2,
	}
	dh := &PrimsEdge{
		u:      "d",
		v:      "h",
		weight: 4.1,
	}
	dg := &PrimsEdge{
		u:      "d",
		v:      "g",
		weight: 7.2,
	}

	ef := &PrimsEdge{
		u:      "e",
		v:      "f",
		weight: 6,
	}
	fh := &PrimsEdge{
		u:      "f",
		v:      "h",
		weight: 2.1,
	}

	hi := &PrimsEdge{
		u:      "h",
		v:      "i",
		weight: 10,
	}
	gh := &PrimsEdge{
		u:      "g",
		v:      "h",
		weight: 14,
	}
	gi := &PrimsEdge{
		u:      "g",
		v:      "i",
		weight: 9,
	}
	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	actual := KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, bc, bd, cf, ce, de, dh, dg, ef, fh, hi, gi, gh})
	expected := []*PrimsEdge{ab, ac, bc, cf, fh, dh, dg, gi}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
