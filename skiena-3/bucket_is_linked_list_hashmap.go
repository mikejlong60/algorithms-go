package skiena_3

import (
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/option"
	"math"
)

type HashMap2[K, V any] struct {
	underlying []*linked_list.LinkedList[KeyValuePair[K, V]]
	eq         func(k1, k2 KeyValuePair[K, V]) bool
	hash       func(k K) uint32
}

func New2[K, V any](eq func(k1, k2 KeyValuePair[K, V]) bool, hash func(k K) uint32, capacity int32) HashMap2[K, V] {
	return HashMap2[K, V]{
		eq:         eq,
		hash:       hash,
		underlying: make([]*linked_list.LinkedList[KeyValuePair[K, V]], capacity),
	}
}

func Get2[K, V any](m HashMap2[K, V], k K, p func(xs KeyValuePair[K, V]) bool) option.Option[KeyValuePair[K, V]] {
	a := m.hash(k)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying))))

	b := m.underlying[idx]
	c := linked_list.Filter[KeyValuePair[K, V]](b, p)

	if c != nil {
		return option.Some[KeyValuePair[K, V]]{linked_list.Head[KeyValuePair[K, V]](c)}
	} else {
		return option.None[KeyValuePair[K, V]]{}
	}
}

func Delete2[K, V any](m HashMap2[K, V], k K, p func(xs KeyValuePair[K, V]) bool) HashMap2[K, V] {
	a := m.hash(k)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying))))
	b := m.underlying[idx]
	notP := func(xs KeyValuePair[K, V]) bool {
		return !p(xs)
	}
	c := linked_list.Filter[KeyValuePair[K, V]](b, notP)
	m.underlying[idx] = c
	return m
}

func replace[K, V any](m HashMap2[K, V], kv KeyValuePair[K, V], p func(xs KeyValuePair[K, V]) bool, idx int32) *linked_list.LinkedList[KeyValuePair[K, V]] {
	b := m.underlying[idx]
	c := linked_list.Filter[KeyValuePair[K, V]](b, p)

	if c != nil {
		if c.Tail != nil {
			c = linked_list.DeleteInBigOh1(c)
			c = linked_list.Push(kv, c)
		} else {
			c.Head = kv
			c.Tail = nil
		}
		return c
	} else {
		b = linked_list.Push[KeyValuePair[K, V]](kv, b)
		return b
	}
}

// Sets or replaces the current value for the key in BigOH(n) worst case.
func Set2[K, V any](m HashMap2[K, V], kv KeyValuePair[K, V], p func(xs KeyValuePair[K, V]) bool) HashMap2[K, V] { //*linked_list.LinkedList[KeyValuePair[K, V]] {
	a := m.hash(kv.key)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying))))
	putKeyValuePair2InRightBucket := func() *linked_list.LinkedList[KeyValuePair[K, V]] {
		return replace(m, kv, p, idx)
	}
	d := putKeyValuePair2InRightBucket()
	m.underlying[idx] = d
	return m
}
