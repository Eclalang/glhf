package glhf

import "github.com/go-gl/gl/v4.4-core/gl"

// Init initializes OpenGL by loading function pointers from the active OpenGL context.
// This function must be manually run inside the main thread (using "github.com/faiface/mainthread"
// package).
//
// It must be called under the presence of an active OpenGL context, e.g., always after calling
// window.MakeContextCurrent(). Also, always call this function when switching contexts.
func Init() {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	gl.Enable(gl.BLEND)
	gl.Enable(gl.SCISSOR_TEST)
	gl.BlendEquation(gl.FUNC_ADD)
}

