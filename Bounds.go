package glhf

import "github.com/go-gl/gl/v4.4-core/gl"

// Bounds sets the drawing bounds in pixels. Drawing outside bounds is always discarted.
//
// Calling this function is equivalent to setting viewport and scissor in OpenGL.
func Bounds(x, y, w, h int) {
	gl.Viewport(int32(x), int32(y), int32(w), int32(h))
	gl.Scissor(int32(x), int32(y), int32(w), int32(h))
}
