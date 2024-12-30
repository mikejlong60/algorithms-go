package skiena_3

import (
	"bitbucket.org/pcastools/hash"
	"fmt"
	"github.com/greymatter-io/golangz/option"
	"testing"
)

func TestGoHashMap(t *testing.T) {
	s := map[int]any{3: nil, 2: nil, 1: nil}
	s[12] = nil
	_, ok := s[12]
	if !ok {
		t.Errorf("Key 12 not added")
	}
	delete(s, 12)
	_, ok = s[12]
	if ok {
		t.Errorf("Key 12 not deleted")
	}
}

func TestYourOwnHashMap(t *testing.T) {
	fFNV32a := func(x int32) uint32 {
		return hash.Int32(x)
	}
	eq := func(k1, k2 int32) bool {
		if k1 == k2 {
			return true
		} else {
			return false
		}
	}
	m := New[int32, string](KeyValuePair[int32, string]{-1, ""}, eq, fFNV32a)
	k := int32(1234234)
	m = Set(m, KeyValuePair[int32, string]{k, "fred"})
	//f := func(x string) string {
	//	return x
	//}
	//option.Map(Get(m, k), f)
	err := option.GetOrElse(Get(m, int32(188)), fmt.Sprintf("Should not have found:%v in HashMap", 188))
	if err != fmt.Sprintf("Should not have found:%v in HashMap", 188) {
		t.Errorf(err)
	}
}
