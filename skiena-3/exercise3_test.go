package skiena_3

import (
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestReverseLinkedList(t *testing.T) {
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	ge := propcheck.ChooseArray(5, 20, propcheck.String(40))

	reverseList := func(xss []string) *linked_list.LinkedList[string] {
		var l *linked_list.LinkedList[string]
		var i = 0
		for {
			if len(xss) == 0 {
				break
			}
			l = linked_list.Push(xss[i], l)
			if i+1 == len(xss) {
				break
			} else {
				i++
			}
		}
		return l
	}

	prop := propcheck.ForAll(ge,
		"Validate reversal of singly-linked list in liner time big Oh(N)  \n",
		func(xs []string) []string {
			return xs
		},
		func(xss []string) (bool, error) {
			var errors error

			l := reverseList(xss)

			if xss[len(xss)-1] != linked_list.Head(l) {
				errors = multierror.Append(errors, fmt.Errorf("Actual: %v Expected:%v", linked_list.Head(l), xss[len(xss)-1]))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]string](t, result)
}
