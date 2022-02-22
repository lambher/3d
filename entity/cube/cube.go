package cube

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Cube struct {
	Position      mgl32.Vec3
	rotationX     float32
	rotationY     float32
	textureSide   uint32
	textureTop    uint32
	textureBottom uint32
}

func (c Cube) Draw() {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(c.Position.X(), c.Position.Y(), c.Position.Z())
	gl.Rotatef(c.rotationX, 1, 0, 0)
	gl.Rotatef(c.rotationY, 0, 1, 0)

	gl.Color4f(1, 1, 1, 1)

	c.drawBottom()
	c.drawTop()
	c.drawSides()
}

func (c Cube) drawBottom() {
	gl.BindTexture(gl.TEXTURE_2D, c.textureBottom)

	gl.Begin(gl.QUADS)

	gl.Normal3f(0, -1, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, 1)

	gl.End()
}

func (c Cube) drawSides() {
	gl.BindTexture(gl.TEXTURE_2D, c.textureSide)
	gl.Begin(gl.QUADS)

	gl.Normal3f(0, 0, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, -1, -1)

	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, 1)

	gl.Normal3f(1, 0, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, -1, 1)

	gl.Normal3f(-1, 0, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, -1)

	gl.End()
}
func (c Cube) drawTop() {
	gl.BindTexture(gl.TEXTURE_2D, c.textureTop)
	gl.Begin(gl.QUADS)

	gl.Normal3f(0, 1, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, -1)

	gl.End()
}
