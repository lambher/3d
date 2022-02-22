package scene

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/lambher/3d/entity"
	"github.com/lambher/3d/texture"
)

type Scene struct {
	t        *texture.Texture
	width    int
	height   int
	Entities []entity.Entity
}

func (s *Scene) Update() {
	for _, e := range s.Entities {
		e.Update()
	}
}

func (s Scene) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	for _, e := range s.Entities {
		e.Draw()
	}
}
