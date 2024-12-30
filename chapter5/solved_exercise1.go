package chapter5

import (
	"fmt"
	"sync/atomic"
)

func PeakOfAList[A any](xs []A, processorPred func([]A) []A) []A {
	r := processorPred(xs)
	if len(r) > 0 {
		return r
	} else {
		//slice the array in half and send it off recursively
		i := len(xs) / 2
		left := xs[0:i]
		right := xs[i:]

		a := PeakOfAList(left, processorPred)
		b := PeakOfAList(right, processorPred)
		return PeakOfAList(append(a, b...), processorPred)
	}
}

var currentGoRoutines uint64

func async[A any](xs []A, processorPred func([]A) []A, c chan []A) {
	go func() {
		actual := processorPred(xs)
		fmt.Print(actual)
		c <- actual
	}()
}

func PeakOfAList2[A any](xs []A, processorPred func([]A) []A) []A {
	r := processorPred(xs)
	if len(r) > 0 {
		return r
	} else {
		//slice the array in half and send it off recursively
		i := len(xs) / 2
		left := xs[0:i]
		right := xs[i:]

		if atomic.LoadUint64(&currentGoRoutines) < 10 { //TODO fix thus byug
			fmt.Println("spawning a new go routine")
			atomic.AddUint64(&currentGoRoutines, 1)

			c := make(chan []A, 2)
			async(left, processorPred, c)
			async(right, processorPred, c)

			ll := <-c
			rr := <-c

			return PeakOfAList2(append(ll, rr...), processorPred)
		} else {
			fmt.Println("NOT spawning a new go routine")
			ll := PeakOfAList2(left, processorPred)
			rr := PeakOfAList2(right, processorPred)
			return PeakOfAList2(append(ll, rr...), processorPred)
		}
	}
}
