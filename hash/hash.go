package uhash

import (
	"hash"
)

//	Convenience short-hand for `h.Write(data)`, then `h.Sum(b)`.
func WriteAndSum(h hash.Hash, data, b []byte) (sum []byte, err error) {
	if _, err = h.Write(data); err == nil {
		sum = h.Sum(b)
	}
	return
}

//	Fowler/Noll/Vo '1'
func Fnv1(vals []int) (h int) {
	var ui uint64 = 14695981039346656037
	h = int(ui)
	for i := 0; i < len(vals); i++ {
		h = (h * 16777619) ^ vals[i]
	}
	return
}

//	Fowler/Noll/Vo '1a'
func Fnv1a(vals []int) (h int) {
	var ui uint64 = 14695981039346656037
	h = int(ui)
	for i := 0; i < len(vals); i++ {
		h = (h ^ vals[i]) * 16777619
	}
	return
}

//	A minor update to Bernstein's hash replaces addition with XOR for the combining step.
func ModifiedBernstein(vals []int) (h int) {
	h = 0
	for i := 0; i < len(vals); i++ {
		h = 33*h ^ vals[i]
	}
	return
}

//	Bob Jenkins
func OneAtATime(vals []int) (h int) {
	h = 0
	for i := 0; i < len(vals); i++ {
		h += vals[i]
		h += (h << 10)
		h ^= (h >> 6)
	}
	h += (h << 3)
	h ^= (h >> 11)
	h += (h << 15)
	return
}

//	The rotating hash with XOR
func RotatingXor(vals []int) (h int) {
	h = 0
	for i := 0; i < len(vals); i++ {
		h = (h << 4) ^ (h >> 28) ^ vals[i]
	}
	return
}

//	The rotating hash with SUM
func RotatingAdd(vals []int) (h int) {
	h = 0
	for i := 0; i < len(vals); i++ {
		h ^= (h << 5) + (h >> 2) + vals[i]
	}
	return
}
