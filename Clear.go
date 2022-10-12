package glhf

import "github.com/go-gl/gl/v4.4-core/gl"

// Clear clears the current framebuffer or window with the given color.
func Clear(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
