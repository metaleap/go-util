package glutil

import (
	"fmt"
	"log"
	"strings"

	gl "github.com/chsc/gogl/gl42"
)

type TShaderProgram struct {
	Name string
	Program gl.Uint
	AttrLocs map[string]gl.Uint
	UnifLocs map[string]gl.Int
}

	func NewShaderProgram (name string, glCShader, glFShader, glGShader, glTcShader, glTeShader, glVShader gl.Uint) (*TShaderProgram, error) {
		var err error
		var hasAtt = false
		var glProg = gl.CreateProgram()
		var glStatus gl.Int
		if glCShader != 0 { hasAtt = true; gl.AttachShader(glProg, glCShader) }
		if glFShader != 0 { hasAtt = true; gl.AttachShader(glProg, glFShader) }
		if glGShader != 0 { hasAtt = true; gl.AttachShader(glProg, glGShader) }
		if glTcShader != 0 { hasAtt = true; gl.AttachShader(glProg, glTcShader) }
		if glTeShader != 0 { hasAtt = true; gl.AttachShader(glProg, glTeShader) }
		if glVShader != 0 { hasAtt = true; gl.AttachShader(glProg, glVShader) }
		if (!hasAtt) {
			gl.DeleteProgram(glProg)
			return nil, fmt.Errorf("No shader attachments specified for program %s", name)
		}
		gl.LinkProgram(glProg)
		if glCShader != 0 { gl.DetachShader(glProg, glCShader); gl.DeleteShader(glCShader) }
		if glFShader != 0 { gl.DetachShader(glProg, glFShader); gl.DeleteShader(glFShader) }
		if glGShader != 0 { gl.DetachShader(glProg, glGShader); gl.DeleteShader(glGShader) }
		if glTcShader != 0 { gl.DetachShader(glProg, glTcShader); gl.DeleteShader(glTcShader) }
		if glTeShader != 0 { gl.DetachShader(glProg, glTeShader); gl.DeleteShader(glTeShader) }
		if glVShader != 0 { gl.DetachShader(glProg, glVShader); gl.DeleteShader(glVShader) }
		if gl.GetProgramiv(glProg, gl.LINK_STATUS, &glStatus); glStatus == 0 {
			err = fmt.Errorf("SHADER PROGRAM %s: %s", name, ShaderInfoLog(glProg, false))
		}
		return &TShaderProgram { name, glProg, map[string]gl.Uint {}, map[string]gl.Int {} }, err
	}

	func (me *TShaderProgram) CleanUp () {
		gl.DeleteProgram(me.Program)
	}

	func (me *TShaderProgram) HasAttr (name string) bool {
		return ShaderIsAttrLocation(me.AttrLocs[name])
	}

	func (me *TShaderProgram) HasUnif (name string) bool {
		return ShaderIsUnifLocation(me.UnifLocs[name])
	}

	func (me *TShaderProgram) SetAttrLocations (attribNames ... string) {
		var loc gl.Uint
		for _, attribName := range attribNames {
			loc = ShaderLocationA(me.Program, attribName)
			if me.AttrLocs[attribName] = loc; ShaderIsAttrLocation(loc) {
				gl.EnableVertexAttribArray(me.AttrLocs[attribName])
			}
			LastError("")
		}
	}

	func (me *TShaderProgram) SetUnifLocations (unifNames ...string) {
		for _, unifName := range unifNames {
			me.UnifLocs[unifName] = ShaderLocationU(me.Program, unifName)
			LastError("")
		}
	}

func ShaderInfoLog (shaderOrProgram gl.Uint, isShader bool) string {
	var err error
	var l = gl.Sizei(256)
	var s = gl.GLStringAlloc(l)
	defer gl.GLStringFree(s)
	if isShader { gl.GetShaderInfoLog(shaderOrProgram, l, nil, s) } else { gl.GetProgramInfoLog(shaderOrProgram, l, nil, s) }
	if err = LastError("ShaderInfoLog(s=%v)", isShader); err != nil { return fmt.Sprintf("gl.ShaderInfoLog(s=%v) call failed: %v", isShader, err.Error()) }
	return gl.GoString(s)
}

func ShaderIsAttrLocation (loc gl.Uint) bool {
	return (loc >= 0) && (loc < 4294967295)
}

func ShaderIsUnifLocation (loc gl.Int) bool {
	return loc >= 0
}

func ShaderLocation (glProg gl.Uint, name string, isAtt bool) gl.Int {
	var loc gl.Int
	var s = gl.GLString(name)
	defer gl.GLStringFree(s)
	if isAtt { loc = gl.GetAttribLocation(glProg, s) } else { loc = gl.GetUniformLocation(glProg, s) }
	return loc
}

func ShaderLocationA (glProg gl.Uint, name string) gl.Uint {
	return gl.Uint(ShaderLocation(glProg, name, true))
}

func ShaderLocationU (glProg gl.Uint, name string) gl.Int {
	return ShaderLocation(glProg, name, false)
}

func ShaderSource (name string, shader gl.Uint, source string, defines map[string]interface{}, logPrint bool, glslVersion string) error {
	var src []*gl.Char
	var i, l = 1, len(defines)
	var lines = make([]string, (l * 5) + 3)
	var joined string
	lines[0] = "#version " + glslVersion + " core\n"
	for dk, dv := range defines {
		lines[i + 0] = "#define "
		lines[i + 1] = dk
		lines[i + 2] = " "
		lines[i + 3] = fmt.Sprintf("%v", dv)
		lines[i + 4] = "\n"
		i = i + 5
	}
	lines[i] = "#line 1\n"
	lines[i + 1] = source
	joined = strings.Join(lines, "")
	src = gl.GLStringArray(lines ...)
	defer gl.GLStringArrayFree(src)
	gl.ShaderSource(shader, gl.Sizei(len(src)), &src[0], nil)
	if logPrint { log.Printf("\n\n------------------------------\n%s\n\n", joined) }
	return LastError("ShaderSource(name=%v source=%v)", name, joined)
}
