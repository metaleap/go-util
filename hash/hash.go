package uhash

import (
	"encoding/base64"
	"hash"
)

//	Returns the `enc`'s string-encoding of the specified `Hash` for `data`.
//
//	If `enc` is `nil`, then `base64.URLEncoding` will be used instead.
func EncodeToString(h hash.Hash, data []byte, enc *base64.Encoding) string {
	h.Write(data)
	if enc == nil {
		enc = base64.URLEncoding
	}
	return enc.EncodeToString(h.Sum(nil))
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
