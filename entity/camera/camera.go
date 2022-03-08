package camera

import (
	"math"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lambher/3d/context"
	"github.com/lambher/3d/entity"
	"github.com/lambher/3d/input"
	"github.com/lambher/3d/models"
)

type Camera struct {
	ctx *context.Context

	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Width    int
	Height   int
}

func NewCamera(ctx *context.Context, camera models.Camera) *Camera {
	return &Camera{
		ctx:      ctx,
		Position: entity.Vec3(camera.Position),
		Width:    int(camera.View.X),
		Height:   int(camera.View.Y),
	}
}

func (c Camera) Draw() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	f := float64(c.Width)/float64(c.Height) - 1
	gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 100.0)
	gl.Rotatef(c.Rotation.X(), 1, 0, 0)
	gl.Rotatef(c.Rotation.Y(), 0, 1, 0)
	gl.Rotatef(c.Rotation.Z(), 0, 0, 1)
	gl.Translatef(c.Position.X(), c.Position.Y(), c.Position.Z())
}

func (c *Camera) Update() {
	if c.ctx.GetInputValue(input.MoveForward) {
		c.Position = c.Position.Add(rotate(mgl32.Vec3{0, 0, 0.1}, c.Rotation))
	}

	if c.ctx.GetInputValue(input.MoveBackward) {
		c.Position = c.Position.Add(rotate(mgl32.Vec3{0, 0, -0.1}, c.Rotation))
	}

	if c.ctx.GetInputValue(input.MoveLeft) {
		c.Position = c.Position.Add(rotate(mgl32.Vec3{0.1, 0, 0}, c.Rotation))
	}
	if c.ctx.GetInputValue(input.MoveRight) {
		c.Position = c.Position.Add(rotate(mgl32.Vec3{-0.1, 0, 0}, c.Rotation))
	}

	if c.ctx.GetInputValue(input.MoveUp) {
		c.Position = c.Position.Add(mgl32.Vec3{0, .1, 0})
	}
	if c.ctx.GetInputValue(input.MoveDown) {
		c.Position = c.Position.Add(mgl32.Vec3{0, -.1, 0})
	}

	if c.ctx.GetInputValue(input.LookLeft) {
		c.Rotation = c.Rotation.Add(mgl32.Vec3{0, -1, 0})
	}

	if c.ctx.GetInputValue(input.LookRight) {
		c.Rotation = c.Rotation.Add(mgl32.Vec3{0, 1, 0})
	}

	if c.ctx.GetInputValue(input.LookUp) {
		c.Rotation = c.Rotation.Add(mgl32.Vec3{-1, 0, 0})
	}
	if c.ctx.GetInputValue(input.LookDown) {
		c.Rotation = c.Rotation.Add(mgl32.Vec3{1, 0, 0})
	}

	//c.Position = c.Position.Add(mgl32.Vec3{0, 0, -0.01})
	//c.rotationX += 0.5
}

// https://stackoverflow.com/questions/14607640/rotating-a-vector-in-3d-space
func rotate(vec mgl32.Vec3, rotation mgl32.Vec3) mgl32.Vec3 {
	//fmt.Println(rotation.Z())
	//xPrim := vec.X()*cos(rotation.Z()) - vec.Y()*sin(rotation.Z())
	//yPrim := vec.X()*sin(rotation.Z()) + vec.Y()*cos(rotation.Z())
	//zPrim := vec.Z()

	// rotationY
	teta := mgl32.DegToRad(rotation.Y())
	xPrim := vec.X()*cos(teta) - vec.Z()*sin(teta)
	yPrim := vec.Y()
	zPrim := -vec.X()*sin(teta) + vec.Z()*cos(teta)

	vec = mgl32.Vec3{
		xPrim,
		yPrim,
		zPrim,
	}

	// rotationX
	teta = mgl32.DegToRad(-rotation.X())
	xPrim = vec.X()
	yPrim = vec.Y()*cos(teta) - vec.Z()*sin(teta)
	zPrim = vec.Y()*sin(teta) + vec.Z()*cos(teta)

	return mgl32.Vec3{
		xPrim,
		yPrim,
		zPrim,
	}
}

func cos(float322 float32) float32 {
	return float32(math.Cos(float64(float322)))
}

func sin(float322 float32) float32 {
	return float32(math.Sin(float64(float322)))
}
