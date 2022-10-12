package glhf

import "github.com/go-gl/gl/v4.4-core/gl"

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}
