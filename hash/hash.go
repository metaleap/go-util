//	http://eternallyconfuzzled.com/tuts/algorithms/jsw_tut_hashing.aspx
package uhash

func Fnv1(ints []int) (h int) {
	var ui uint64 = 14695981039346656037
	h = int(ui)
	for _, v := range ints {
		h = (h * 16777619) ^ v
	}
	return
}

func Fnv1a(ints []int) (h int) {
	var ui uint64 = 14695981039346656037
	h = int(ui)
	for _, v := range ints {
		h = (h ^ v) * 16777619
	}
	return
}

func ModifiedBernstein(ints []int) (h int) {
	h = 0
	for _, v := range ints {
		h = 33*h ^ v
	}
	return
}

func OneAtATime(ints []int) (h int) {
	h = 0
	for _, v := range ints {
		h += v
		h += (h << 10)
		h ^= (h >> 6)
	}
	h += (h << 3)
	h ^= (h >> 11)
	h += (h << 15)
	return
}

func RotatingXor(ints []int) (h int) {
	h = 0
	for _, v := range ints {
		h = (h << 4) ^ (h >> 28) ^ v
	}
	return
}

func RotatingAdd(ints []int) (h int) {
	h = 0
	for _, v := range ints {
		h ^= (h << 5) + (h >> 2) + v
	}
	return
}
