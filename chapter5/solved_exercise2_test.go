package chapter5

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

var eq = func(l, r DayStockPrice) bool {
	if l.day == r.day {
		return true
	} else {
		return false
	}
}

func TestMostProfitAscending2(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}

	actual := MostProfit([]DayStockPrice{a, b})
	expected := []DayStockPrice{a, b}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitDescending2(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 9,
	}

	actual := MostProfit([]DayStockPrice{a, b})
	expected := []DayStockPrice{}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitAscending4(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}
	c := DayStockPrice{
		day:   2,
		price: 12,
	}
	d := DayStockPrice{
		day:   3,
		price: 13,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d})
	expected := []DayStockPrice{a, d}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitDescending4(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 13,
	}
	b := DayStockPrice{
		day:   1,
		price: 12,
	}

	c := DayStockPrice{
		day:   2,
		price: 11,
	}
	d := DayStockPrice{
		day:   3,
		price: 10,
	}
	actual := MostProfit([]DayStockPrice{a, b, c, d})
	expected := []DayStockPrice{}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitAscending8(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}
	c := DayStockPrice{
		day:   2,
		price: 12,
	}
	d := DayStockPrice{
		day:   3,
		price: 13,
	}

	e := DayStockPrice{
		day:   4,
		price: 14,
	}
	f := DayStockPrice{
		day:   5,
		price: 15,
	}
	g := DayStockPrice{
		day:   6,
		price: 16,
	}
	h := DayStockPrice{
		day:   7,
		price: 17,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f, g, h})
	expected := []DayStockPrice{a, h}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfitDescending8(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 20,
	}
	b := DayStockPrice{
		day:   1,
		price: 19,
	}
	c := DayStockPrice{
		day:   2,
		price: 18,
	}
	d := DayStockPrice{
		day:   3,
		price: 17,
	}

	e := DayStockPrice{
		day:   4,
		price: 16,
	}
	f := DayStockPrice{
		day:   5,
		price: 15,
	}
	g := DayStockPrice{
		day:   6,
		price: 14,
	}
	h := DayStockPrice{
		day:   7,
		price: 13,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f, g, h})
	expected := []DayStockPrice{}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfit8PeakRight(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}
	c := DayStockPrice{
		day:   2,
		price: 12,
	}
	d := DayStockPrice{
		day:   3,
		price: 13,
	}

	e := DayStockPrice{
		day:   4,
		price: 140,
	}
	f := DayStockPrice{
		day:   5,
		price: 15,
	}
	g := DayStockPrice{
		day:   6,
		price: 16,
	}
	h := DayStockPrice{
		day:   7,
		price: 17,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f, g, h})
	expected := []DayStockPrice{a, e}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfit8PeakLeft(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}
	c := DayStockPrice{
		day:   2,
		price: 12,
	}
	d := DayStockPrice{
		day:   3,
		price: 130,
	}

	e := DayStockPrice{
		day:   4,
		price: 14,
	}
	f := DayStockPrice{
		day:   5,
		price: 15,
	}
	g := DayStockPrice{
		day:   6,
		price: 16,
	}
	h := DayStockPrice{
		day:   7,
		price: 17,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f, g, h})
	expected := []DayStockPrice{a, d}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfit8PeakLeft2(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 11,
	}
	c := DayStockPrice{
		day:   2,
		price: 120,
	}
	d := DayStockPrice{
		day:   3,
		price: 13,
	}

	e := DayStockPrice{
		day:   4,
		price: 14,
	}
	f := DayStockPrice{
		day:   5,
		price: 15,
	}
	g := DayStockPrice{
		day:   6,
		price: 16,
	}
	h := DayStockPrice{
		day:   7,
		price: 17,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f, g, h})
	expected := []DayStockPrice{a, c}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestMostProfit8SolutionSplitAcrossHalves(t *testing.T) {
	a := DayStockPrice{
		day:   0,
		price: 10,
	}
	b := DayStockPrice{
		day:   1,
		price: 3,
	}
	c := DayStockPrice{
		day:   2,
		price: 12,
	}
	d := DayStockPrice{
		day:   3,
		price: 13,
	}

	e := DayStockPrice{
		day:   4,
		price: 14,
	}
	f := DayStockPrice{
		day:   5,
		price: 15,
	}
	g := DayStockPrice{
		day:   6,
		price: 160,
	}
	h := DayStockPrice{
		day:   7,
		price: 17,
	}

	actual := MostProfit([]DayStockPrice{a, b, c, d, e, f, g, h})
	expected := []DayStockPrice{b, g}

	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}
