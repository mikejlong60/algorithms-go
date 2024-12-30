package skiena_3

import (
	"bufio"
	"fmt"
	"github.com/greymatter-io/golangz/option"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"os"
	"strings"
	"testing"
	"time"
)

func TestBalancedBinaryTreeFound(t *testing.T) {

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
	g0 := propcheck.ChooseArray(100000, 100000, propcheck.ChooseInt(-1000000, 1000000))

	fg := func(fx []int) func(rng propcheck.SimpleRNG) (propcheck.Pair[int, []int], propcheck.SimpleRNG) {
		a := propcheck.ChooseInt(0, len(fx))
		g := func(x int) propcheck.Pair[int, []int] {
			return propcheck.Pair[int, []int]{x, fx}
		}
		r := propcheck.Map(a, g)
		return r
	}

	g1 := propcheck.FlatMap(g0, fg)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g1,
		"Make a binary tree",
		func(a propcheck.Pair[int, []int]) propcheck.Pair[int, []int] {
			return a
		},
		func(a propcheck.Pair[int, []int]) (bool, error) {

			btree := BinaryTree(a.B, lt)
			var errors error
			if len(a.B) > 0 {
				idx := a.A
				f := func(x int) int {
					if x != a.B[idx] {
						errors = multierror.Append(errors, fmt.Errorf("Not Found"))
					}
					return x
				}
				option.Map(Find(btree, a.B[idx], lt, eq), f)
			}

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[int, []int]](t, result)
}

func TestBalancedBinaryTreeNotFound(t *testing.T) {

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
	g0 := propcheck.ChooseArray(0, 62, propcheck.ChooseInt(0, 100))

	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Make a binary tree",
		func(a []int) []int {
			return a
		},
		func(a []int) (bool, error) {
			btree := BinaryTree(a, lt)
			var errors error
			if len(a) > 0 {
				f := func(x int) int {
					errors = fmt.Errorf("Should not have Found:%v", x)
					return x
				}
				option.Map(Find(btree, -10, lt, eq), f) //-10 is not in tree
			}

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestPhonePad(t *testing.T) {
	lt := func(l, r string) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r string) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}

	readLinesFromFile := func(filename string) ([]string, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		lines := []string{}
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			word := scanner.Text()
			if len(word) == 3 {
				lines = append(lines, strings.ToLower(word))
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return lines, nil
	}

	loadTreeWithWords := func(filename string) (*Node[string], error) {
		a, err := readLinesFromFile(filename)
		r := sets.ToSet(a, lt, eq)
		btree := BinaryTree(r, lt)
		return btree, err
	}

	getAllWordsFromThreeNumbers := func(numbers []int, keyPadDict map[int][]string, wordDict *Node[string]) []string {
		wordCandidates := []string{}
		firstLetters := keyPadDict[numbers[0]]
		secondLetters := keyPadDict[numbers[1]]
		thirdLetters := keyPadDict[numbers[2]]
		for _, first := range firstLetters {
			for _, second := range secondLetters {
				for _, third := range thirdLetters {
					threeLetterWord := fmt.Sprintf("%v%v%v", first, second, third)
					wordCandidates = append(wordCandidates, strings.ToLower(threeLetterWord))
				}
			}
		}

		foundWords := []string{}
		f := func(w string) string {
			foundWords = append(foundWords, w)
			return w
		}
		for _, word := range wordCandidates {
			option.Map(Find(wordDict, word, lt, eq), f)
		}
		return foundWords
	}

	wordDict, _ := loadTreeWithWords("words.txt")

	keyPadDict := map[int][]string{
		1: {"A", "B", "C"},
		2: {"D", "E", "F"},
		3: {"G", "H", "I"},
		4: {"J", "K", "L"},
		5: {"M", "N", "O"},
		6: {"P", "Q", "R"},
		7: {"S", "T", "U"},
		8: {"V", "W", "X"},
		0: {"Y", "Z"},
	}

	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	ge := propcheck.ChooseArray(3, 3, propcheck.ChooseInt(0, 9))

	prop := propcheck.ForAll(ge,
		"Validate phone pad word  \n",
		func(xs []int) []int {
			return xs
		},
		func(xss []int) (bool, error) {
			var errors error

			l := getAllWordsFromThreeNumbers(xss, keyPadDict, wordDict)
			for _, word := range l {
				err := option.GetOrElse(Find(wordDict, word, lt, eq), fmt.Sprintf("Word:%v not in dictionary", word))
				if err == fmt.Sprintf("Word:%v not in dictionary", word) {
					errors = multierror.Append(errors, fmt.Errorf(err))
				}
			}

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
