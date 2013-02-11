package ustr

import (
	"bytes"
	"fmt"
)

//	A convenience wrapper for bytes.Buffer.
type Buffer struct {
	bytes.Buffer
}

//	Short-hand for bytes.Buffer.WriteString(fmt.Sprintf(format, fmtArgs...))
func (me *Buffer) Write(format string, fmtArgs ...interface{}) {
	if len(fmtArgs) > 0 {
		format = fmt.Sprintf(format, fmtArgs...)
	}
	me.Buffer.WriteString(format)
}

//	Short-hand for bytes.Buffer.WriteString(fmt.Sprintf(format+"\n", fmtArgs...))
func (me *Buffer) Writeln(format string, fmtArgs ...interface{}) {
	me.Write(format+"\n", fmtArgs...)
}
