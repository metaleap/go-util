//	http://eternallyconfuzzled.com/tuts/algorithms/jsw_tut_hashing.aspx
package uhash

func Fnv1(vals []int) (h int) {
	var ui uint64 = 14695981039346656037
	h = int(ui)
	for i := 0; i < len(vals); i++ {
		h = (h * 16777619) ^ vals[i]
	}
	return
}

func Fnv1a(vals []int) (h int) {
	var ui uint64 = 14695981039346656037
	h = int(ui)
	for i := 0; i < len(vals); i++ {
		h = (h ^ vals[i]) * 16777619
	}
	return
}

func ModifiedBernstein(vals []int) (h int) {
	h = 0
	for i := 0; i < len(vals); i++ {
		h = 33*h ^ vals[i]
	}
	return
}

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

func RotatingXor(vals []int) (h int) {
	h = 0
	for i := 0; i < len(vals); i++ {
		h = (h << 4) ^ (h >> 28) ^ vals[i]
	}
	return
}

func RotatingAdd(vals []int) (h int) {
	h = 0
	for i := 0; i < len(vals); i++ {
		h ^= (h << 5) + (h >> 2) + vals[i]
	}
	return
}
