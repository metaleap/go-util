package ugo

//	Appends v to sl only if sl does not already contain v.
func AppendUniqueInt(sl *[]int, v int) {
	for _, slv := range *sl {
		if slv == v {
			return
		}
	}
	*sl = append(*sl, v)
}
