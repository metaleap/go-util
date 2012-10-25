package glutil

import (
	"errors"
	"fmt"
	"strings"

	gl "github.com/chsc/gogl/gl42"

	strutil "github.com/go3d/go-util/str"
	util "github.com/go3d/go-util"
)

var (
	Version = [2]int { 0, 0 }

	extensionPrefixes = []string { "GL_ARB_", "GL_ATI_", "GL_S3_", "GL_EXT_", "GL_IBM_", "GL_KTX_", "GL_NV_", "GL_NVX_", "GL_OES_", "GL_SGIS_", "GL_SGIX_", "GL_SUN_", "GL_APPLE_" }
	extensions []string = nil
	IsGl32, IsGl33, IsGl40, IsGl41, IsGl42, IsGl43 bool
	GlSlVersion = "150"
)

func Extension (name string) bool {
	if strings.HasPrefix(name, "GL_") { return strutil.InSliceAt(Extensions(), name) >= 0 }
	for _, ep := range extensionPrefixes { if strutil.InSliceAt(Extensions(), ep + name) >= 0 { return true } }
	return false
}

func Extensions () []string {
	if extensions == nil { extensions = strutil.Split(GlStr(gl.EXTENSIONS), " ") }
	return extensions
}

func GlConnInfo () string {
	return fmt.Sprintf("OpenGL %v @ %v %v (GLSL: %v)", GlStr(gl.VERSION), GlStr(gl.VENDOR), GlStr(gl.RENDERER), GlStr(gl.SHADING_LANGUAGE_VERSION))
}

func GlStr (name gl.Enum) string {
	return gl.GoStringUb(gl.GetString(name))
}

func GlVal (name gl.Enum) (gl.Int, error) {
	var ret gl.Int
	gl.GetIntegerv(name, &ret)
	return ret, LastError("Integerv(n=%v)", name)
}

func GlVals (name gl.Enum, num uint) ([]gl.Int, error) {
	var ret = make([]gl.Int, num)
	gl.GetIntegerv(name, &ret[0])
	return ret, LastError("Integervs(n=%v)", name)
}

func LastError (step string, fmtArgs ... interface{}) error {
	var errEnum gl.Enum = gl.GetError()
	var err error
	var ln, errStr string
	if errEnum != 0 {
		if len(fmtArgs) > 0 { step = fmt.Sprintf(step, fmtArgs ...) }
		errStr += fmt.Sprintf("OpenGL error at step '%s':", step)
		switch errEnum {
		case gl.INVALID_ENUM:
			ln = "GL_INVALID_ENUM"
		case gl.INVALID_VALUE:
			ln = "GL_INVALID_VALUE"
		case gl.INVALID_OPERATION:
			ln = "GL_INVALID_OPERATION"
		case gl.OUT_OF_MEMORY:
			ln = "GL_OUT_OF_MEMORY"
		case gl.INVALID_FRAMEBUFFER_OPERATION:
			ln = "GL_INVALID_FRAMEBUFFER_OPERATION"
		default:
			ln = fmt.Sprintf("%v", errEnum)
		}
		errStr += fmt.Sprintf("\t%s", ln)
		err = errors.New(errStr)
	}
	return err
}

func SetVersion () {
	IsGl32, IsGl33, IsGl40, IsGl41, IsGl42, IsGl43 = false, false, false, false, false, false
	GlSlVersion = "120"
	if glVer := util.ParseVersion(GlStr(gl.VERSION)); len(glVer) > 0 {
		if Version[0] = glVer[0]; len(glVer) > 1 { Version[1] = glVer[1] }
		if Version[0] >= 3 {
			if Version[0] == 3 {
				if Version[1] >= 2 { IsGl32, GlSlVersion = true, "150" }
				if Version[1] >= 3 { IsGl33, GlSlVersion = true, "330" }
			}
			if Version[0] == 4 {
				IsGl32, IsGl33, IsGl40, GlSlVersion = true, true, true, "400"
				if Version[1] >= 1 { IsGl41, GlSlVersion = true, "410" }
				if Version[1] >= 2 { IsGl42, GlSlVersion = true, "420" }
				if Version[1] >= 3 { IsGl43, GlSlVersion = true, "430" }
			}
		}
	}
}

func VersionMatch (major, minor int) bool {
	return (Version[0] > major) || ((Version[0] == major) && (Version[1] >= minor))
}
