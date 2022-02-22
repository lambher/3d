package scene

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lambher/3d/entity"
	"github.com/lambher/3d/entity/cube"
	"github.com/lambher/3d/texture"
)

type Basic struct {
	*Scene
}

func NewBasic(t *texture.Texture, width, height int) *Basic {
	return &Basic{
		Scene: &Scene{
			t:      t,
			width:  width,
			height: height,
			Entities: []entity.Entity{
				cube.NewGrass(t, mgl32.Vec3{0, 0, -3}),
				//cube.NewGrass(t, mgl32.Vec3{0, 1, -3}),
			},
		},
	}
}

func (b Basic) Setup() {
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.LIGHTING)

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)

	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	f := float64(b.width)/float64(b.height) - 1
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func (b Basic) Draw() {
	b.Scene.Draw()
}

func (b *Basic) Update() {
	b.Scene.Update()
}
