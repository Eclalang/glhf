package glhf

import "github.com/go-gl/gl/v4.4-core/gl"

// BlendFactor represents a source or destination blend factor.
type BlendFactor int

// Here's the list of all blend factors.
const (
	One              = BlendFactor(gl.ONE)
	Zero             = BlendFactor(gl.ZERO)
	SrcAlpha         = BlendFactor(gl.SRC_ALPHA)
	DstAlpha         = BlendFactor(gl.DST_ALPHA)
	OneMinusSrcAlpha = BlendFactor(gl.ONE_MINUS_SRC_ALPHA)
	OneMinusDstAlpha = BlendFactor(gl.ONE_MINUS_DST_ALPHA)
)

// BlendFunc sets the source and destination blend factor.
func BlendFunc(src, dst BlendFactor) {
	gl.BlendFunc(uint32(src), uint32(dst))
}
