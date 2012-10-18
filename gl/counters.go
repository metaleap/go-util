package glutil

import (
	gl "github.com/chsc/gogl/gl42"
)

func ReadAtomicCounterValues (ac gl.Uint, vals []gl.Uint) error {
	var ptr *gl.Uint
	gl.BindBuffer(gl.ATOMIC_COUNTER_BUFFER, ac)
	for i := 0; i < len(vals); i++ {
		ptr = (*gl.Uint)(gl.MapBufferRange(gl.ATOMIC_COUNTER_BUFFER, OffsetIntPtr(nil, gl.Sizei(i * 4)), SizeOfGlUint, gl.MAP_READ_BIT))
		vals[i] = *ptr
		gl.UnmapBuffer(gl.ATOMIC_COUNTER_BUFFER)
	}
	gl.BindBuffer(gl.ATOMIC_COUNTER_BUFFER, 0)
	return LastError("ReadAtomicCounter(%v)", ac)
}

func ResetAtomicCounters (glPtr gl.Uint, num gl.Sizei, value gl.Uint) {
	var ptr *gl.Uint
	var i gl.Sizei
	gl.BindBuffer(gl.ATOMIC_COUNTER_BUFFER, glPtr)
	for i = 0; i < num; i++ {
		ptr = (*gl.Uint)(gl.MapBufferRange(gl.ATOMIC_COUNTER_BUFFER, OffsetIntPtr(nil, i * 4), SizeOfGlUint, gl.MAP_WRITE_BIT | gl.MAP_INVALIDATE_BUFFER_BIT | gl.MAP_UNSYNCHRONIZED_BIT))
		*ptr = value
		gl.UnmapBuffer(gl.ATOMIC_COUNTER_BUFFER)
	}
	gl.BindBuffer(gl.ATOMIC_COUNTER_BUFFER, 0)
}
