package scene

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/lambher/3d/context"
	"github.com/lambher/3d/entity"
	"github.com/lambher/3d/entity/camera"
	"github.com/lambher/3d/entity/cube"
	"github.com/lambher/3d/models"
)

type Scene struct {
	ctx *context.Context

	Entities []entity.Entity
	Camera   *camera.Camera
	Light    models.Light
}

func (s *Scene) Update() {
	s.Camera.Update()
	for _, e := range s.Entities {
		e.Update()
	}
}

func (s Scene) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	s.Camera.Draw()
	for _, e := range s.Entities {
		e.Draw()
	}
}

func (s *Scene) LoadScene(ctx *context.Context, scene models.Scene) {
	s.Camera = camera.NewCamera(ctx, scene.Camera)
	for _, c := range scene.Cubes {
		s.Entities = append(s.Entities, cube.NewCube(ctx, c))
	}
	s.Light = scene.Light
}
