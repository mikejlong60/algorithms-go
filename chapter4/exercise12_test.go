package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

type Stream struct {
	NumberOfBits int
	Duration     int
}

func makeSchedule(xs []Stream, maxBitsPerSecond int) []Stream {
	var r = []Stream{}
	for a, b := range xs {
		if b.NumberOfBits*b.Duration <= maxBitsPerSecond*b.Duration {
			r = append(r, b)
			r = append(r, xs[0:a]...)
			r = append(r, xs[a+1:]...)
			return r
		}
	}
	return []Stream{}
}

func seq(l, r Stream) bool {
	if l.Duration == r.Duration && l.NumberOfBits == r.NumberOfBits {
		return true
	} else {
		return false
	}
}

func TestMakeValidScheduleWhereFirstStreamFinishesWithinTimeWindow(t *testing.T) {

	//Then you can schedule the rest of the streams in any order

	s1 := Stream{NumberOfBits: 8000, Duration: 1}
	s2 := Stream{NumberOfBits: 7000, Duration: 3}
	s3 := Stream{NumberOfBits: 2000, Duration: 1}
	s4 := Stream{NumberOfBits: 1000, Duration: 2}

	r := 6000

	actual := makeSchedule([]Stream{s1, s2, s3, s4}, r)
	expected := []Stream{s3, s1, s2, s4}
	if !arrays.ArrayEquality(actual, expected, seq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestMakeInvalidScheduleWhereFirstStreamFinishesWithinTimeWindow(t *testing.T) {

	s1 := Stream{NumberOfBits: 8000, Duration: 1}
	s2 := Stream{NumberOfBits: 7000, Duration: 3}
	s3 := Stream{NumberOfBits: 2000, Duration: 1}
	s4 := Stream{NumberOfBits: 1000, Duration: 2}

	r := 500

	actual := makeSchedule([]Stream{s1, s2, s3, s4}, r)
	expected := []Stream{}
	if !arrays.ArrayEquality(actual, expected, seq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
