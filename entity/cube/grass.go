package cube

import (
	"github.com/lambher/3d/context"
	"github.com/lambher/3d/entity"
	"github.com/lambher/3d/models"
	"github.com/lambher/3d/texture"
)

type Grass struct {
	*Cube
}

func NewGrass(ctx *context.Context, cube models.Cube) *Grass {
	return &Grass{
		Cube: &Cube{
			textureSide:   ctx.GetTexture(texture.GrassSide),
			textureTop:    ctx.GetTexture(texture.GrassTop),
			textureBottom: ctx.GetTexture(texture.GrassBottom),
			Position:      entity.Vec3(cube.Position),
		},
	}
}

func (g *Grass) Update() {
	//g.rotationX += 0.1
	//g.rotationY -= 0.5
}

func (g Grass) Draw() {
	g.Cube.Draw()
}
