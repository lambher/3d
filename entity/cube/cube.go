package cube

import "github.com/go-gl/mathgl/mgl32"

type Cube struct {
	Position  mgl32.Vec3
	rotationX float32
	rotationY float32
	texture   uint32
}
