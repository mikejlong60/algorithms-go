package skiena_3

import (
	"fmt"
	"github.com/greymatter-io/golangz/either"
	"testing"
)

type parenPos struct {
	paren string
	pos   int
}

func balancedParenthesesEither(ss string) either.Either[error, int] {
	balanced, errorStarts, maxDistance := balancedParentheses(ss)
	if balanced {
		return either.Right[int]{maxDistance}
	} else {
		return either.Left[error]{fmt.Errorf("Parsing error starts:%v", errorStarts)}
	}
}

// Assumes string is only right and left parentheses
func balancedParentheses(ss string) (bool, int, int) {
	maxDistance := 0
	a := 0
	makeParenPos := func(s string) []parenPos {
		r := []parenPos{}
		for i := 0; i < len(s); i++ {
			r = append(r, parenPos{s[i : i+1], i})
		}
		return r
	}
	s := makeParenPos(ss)
	for {
		if a >= len(s) {
			break
		} else if s[0].paren == ")" {
			break
		} else if s[a].paren == "(" {
			a = a + 1
			continue
		} else { //s[a] == ")"
			m := s[a].pos - s[a-1].pos
			if m > maxDistance {
				maxDistance = m
			}
			sa := s[0 : a-1]
			sb := s[a+1 : len(s)]
			s = append(sa, sb...)
			a = 0
			continue
		}

	}
	if len(s) == 0 {
		return true, -1, maxDistance
	} else {
		return false, s[0].pos, maxDistance
	}
}

func TestEitherChainingWithBind_AKA_FlatMap(t *testing.T) {
	// This example demonstrates the power of monadic(flatmap) chaining.  As soon as you hit an error the computation stops
	// and reports the error without any error handling.  The expression works from the inside out.  I will make it
	// a method on an interface called Either. Then you can write the expression from left to right.

	next := func(maxDistance int) either.Either[error, int] {
		return either.Right[int]{maxDistance}
	}
	next2 := func(maxDistance int) either.Either[error, int] {
		return balancedParenthesesEither("(())")
	}
	next3 := func(maxDistance int) either.Either[error, int] {
		return balancedParenthesesEither("((())")
	}
	next4 := func(maxDistance int) either.Either[error, int] {
		return balancedParenthesesEither("(())")
	}
	start := either.FlatMap(balancedParenthesesEither("((()))"), next)

	cc := either.FlatMap(either.FlatMap(either.FlatMap(either.FlatMap(start, next), next2), next3), next4) //, balancedParenthesesEither("(())"))
	fmt.Println(cc)
}

func TestEitherChainingWithMap(t *testing.T) {
	// This example demonstrates the power of monadic(flatmap) chaining.  As soon as you hit an error the computation stops
	// and reports the error without any error handling.  The expression works from the inside out.  I will make it
	// a method on an interface called Either. Then you can write the expression from left to right.

	next := func(maxDistance int) int {
		return maxDistance + 12
	}
	next2 := func(maxDistance int) either.Either[error, int] {
		return balancedParenthesesEither("(())")
	}
	next3 := func(maxDistance int) either.Either[error, int] {
		return balancedParenthesesEither("((())")
	}
	next4 := func(maxDistance int) either.Either[error, int] {
		return balancedParenthesesEither("(())")
	}

	start := either.Map[error, int](balancedParenthesesEither("((()))"), next)

	cc := either.FlatMap(either.FlatMap(either.FlatMap(either.Map[error, int](start, next), next2), next3), next4)
	fmt.Println(cc)
}

func TestBalancedParenthesesEither(t *testing.T) {
	currentMax := 3
	success := func(maxDistance int) int {
		if maxDistance != currentMax {
			t.Errorf("Actual:%v, Expected:%v", maxDistance, currentMax)
		}
		return maxDistance
	}
	either.Map[error, int](balancedParenthesesEither("(())"), success)
	currentMax = 5
	either.Map[error, int](balancedParenthesesEither("((()))"), success)
	currentMax = 7
	either.Map[error, int](balancedParenthesesEither("((()()))"), success)
	currentMax = 0
	either.Map[error, int](balancedParenthesesEither(""), success)

	//Expected expression parsing failure cases
	actual := either.Map[error, int](balancedParenthesesEither("((())))"), success)
	expectedFailureString := "Parsing error starts:6"
	expectError := func(e either.Either[error, int]) either.Either[error, int] {
		switch v := actual.(type) {
		case either.Left[error]:
			if v.Value.Error() != expectedFailureString {
				t.Errorf("Actual:%v, Expected:%v", v.Value, expectedFailureString)
			}
			return v
		default:
			t.Errorf("Expected Left[error]")
			return v
		}
	}
	expectError(actual)
	expectedFailureString = "Parsing error starts:0"
	actual = either.Map[int, int](balancedParenthesesEither("(((()))"), success)
	expectError(actual)
	actual = either.Map[int, int](balancedParenthesesEither(")())))("), success)
	expectError(actual)
	actual = either.Map[int, int](balancedParenthesesEither("))((()))("), success)
	expectError(actual)
	actual = either.Map[int, int](balancedParenthesesEither(")((()))("), success)
	expectError(actual)
	actual = either.Map[int, int](balancedParenthesesEither(")"), success)
	expectError(actual)
	actual = either.Map[int, int](balancedParenthesesEither(")("), success)
	expectError(actual)
	actual = either.Map[int, int](balancedParenthesesEither("("), success)
	expectError(actual)
}
func TestBalancedParentheses(t *testing.T) {
	actual, startOfError, maxDistance := balancedParentheses("(())")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}
	if maxDistance != 3 {
		t.Errorf("Actual:%v, Expected:%v", maxDistance, 3)
	}
	actual, startOfError, maxDistance = balancedParentheses("((()))")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}
	if maxDistance != 5 {
		t.Errorf("Actual:%v, Expected:%v", maxDistance, 5)
	}
	actual, startOfError, maxDistance = balancedParentheses("((()()))")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}
	if maxDistance != 7 {
		t.Errorf("Actual:%v, Expected:%v", maxDistance, 7)
	}
	actual, startOfError, maxDistance = balancedParentheses("((())))")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 6 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 6)
	}
	actual, startOfError, maxDistance = balancedParentheses("(((()))")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses(")())))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses("))((()))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses(")((()))(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses(")")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses(")(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses("(")
	if actual {
		t.Errorf("Actual:%v, Expected:%v", actual, false)
	}
	if startOfError != 0 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, 0)
	}
	actual, startOfError, maxDistance = balancedParentheses("")
	if !actual {
		t.Errorf("Actual:%v, Expected:%v", actual, true)
	}
	if startOfError != -1 {
		t.Errorf("Actual:%v, Expected:%v", startOfError, -1)
	}
}
