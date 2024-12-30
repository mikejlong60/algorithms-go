package skiena_2

import (
	"fmt"
	"strings"
	"testing"
)

func pangrams(s string) string {
	xs := []byte(strings.ToLower(s))
	alphabet := []byte("abcdefghijklmnopqrstuvwxyz")
	alphabetm := map[byte]struct{}{}
	for _, b := range alphabet {
		alphabetm[b] = struct{}{}
	}
	for _, b := range xs {
		delete(alphabetm, b)
	}
	if len(alphabetm) == 0 {
		return "pangram"
	} else {
		fmt.Println(len(alphabetm))
		return "not pangram"
	}
}

func TestPangram(t *testing.T) {
	actual := pangrams("We promptly judged antique ivory buckles for the next prize")
	if actual != "pangram" {
		t.Errorf("Actual:%v, Expected:%v", actual, "pangram")
	}

	actual = pangrams("We pror the next prize")
	if actual == "pangram" {
		t.Errorf("Actual:%v, Expected:%v", actual, "not pangram")
	}

	actual = pangrams("We promptly judged antique ivory buckles for the prize")
	if actual != "not pangram" {
		t.Errorf("Actual:%v, Expected:%v", actual, "pangram")
	}
}
