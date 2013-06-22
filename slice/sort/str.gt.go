package usort

//#begin-gt -gen.gt N:Str T:string

import "sort"

type sortStr struct {
	descending bool
	slice      []string
}

//	Implements `sort.Interface.Len()`.
func (me *sortStr) Len() int { return len(me.slice) }

//	Implements `sort.Interface.Less()`.
func (me *sortStr) Less(i, j int) bool {
	if me.descending {
		return me.slice[j] < me.slice[i]
	}
	return me.slice[i] < me.slice[j]
}

//	Implements `sort.Interface.Swap()`.
func (me *sortStr) Swap(i, j int) { me.slice[i], me.slice[j] = me.slice[j], me.slice[i] }

//	Returns `sl` sorted by ascending order.
func StrSortAsc(sl []string) []string {
	me := &sortStr{descending: false, slice: sl}
	sort.Sort(me)
	return me.slice
}

//	Returns `sl` sorted by decending order.
func StrSortDesc(sl []string) []string {
	me := &sortStr{descending: true, slice: sl}
	sort.Sort(me)
	return me.slice
}

//#end-gt
