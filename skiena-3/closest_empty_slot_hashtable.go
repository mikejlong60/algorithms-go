package skiena_3

import (
	"github.com/greymatter-io/golangz/option"
	"math"
)

type KeyValuePair[K, V any] struct {
	key   K
	value V
}
type HashMap[K, V any] struct {
	underlying [51]KeyValuePair[K, V]
	zero       KeyValuePair[K, V]
	eq         func(k1, k2 K) bool
	emptyCells int
	hash       func(k K) uint32
}

func New[K, V any](zero KeyValuePair[K, V], eq func(k1, k2 K) bool, hash func(k K) uint32) HashMap[K, V] {

	zeroHashMap := func() [51]KeyValuePair[K, V] {
		a := [51]KeyValuePair[K, V]{}
		for i := 0; i < len(a); i++ {
			a[i] = zero
		}
		return a
	}
	return HashMap[K, V]{
		underlying: zeroHashMap(),
		zero:       zero,
		eq:         eq,
		hash:       hash,
		emptyCells: 10,
	}
}

func Get[K, V any](m HashMap[K, V], k K) option.Option[V] {
	a := m.hash(k)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying))))

	b := m.underlying[idx]
	if !m.eq(b.key, m.zero.key) {
		return option.Some[V]{b.value}
	} else {
		return option.None[V]{}
	}
}

func Set[K, V any](m HashMap[K, V], kv KeyValuePair[K, V]) HashMap[K, V] {
	a := m.hash(kv.key)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying))))

	b := m.underlying[idx]
	assignToHashOrClosestEmpty := func() bool { //Assigns kv pair to either the hashed array index or the closest available spot in the underlying array,
		// moving to the right until you get to the end.
		// Returns true if successful. Otherwise false which indicates array needs resizing.
		if m.eq(b.key, m.zero.key) { //A hash collision. Look forward for an empty spot
			m.underlying[idx] = kv
			m.emptyCells = m.emptyCells - 1
			return true
		} else {
			underlyingSize := int32(len(m.underlying))
			for i := idx; i < underlyingSize; i++ {
				if m.eq(m.underlying[i].key, m.zero.key) {
					m.underlying[i] = kv
					m.emptyCells = m.emptyCells - 1
					return true
				}
			}
			return false
		}
	}

	resizeHashMap := func() HashMap[K, V] {
		return m //TODO Make this really do it, 1. Double the original size. 2. assign the new hash element as part of rehashing the whole old array. 3. Return the new HashMap
	}

	assigned := assignToHashOrClosestEmpty()
	if !assigned {
		resizeHashMap()
		m.underlying[a] = kv
	}

	return m
}
