package cube

import (
	"github.com/lambher/3d/context"
	"github.com/lambher/3d/entity"
	"github.com/lambher/3d/models"
	"github.com/lambher/3d/texture"
)

type Grass2 struct {
	*Cube
}

func NewGrass2(ctx *context.Context, cube models.Cube) *Grass2 {
	return &Grass2{
		Cube: &Cube{
			textureSide:   ctx.GetTexture(texture.GrassSide),
			textureTop:    ctx.GetTexture(texture.GrassTop),
			textureBottom: ctx.GetTexture(texture.GrassBottom),
			Position:      entity.Vec3(cube.Position),
		},
	}
}

func (g *Grass2) Update() {
	//g.rotationX -= 0.1
	//g.rotationY += 0.5
}

func (g Grass2) Draw() {
	g.Cube.Draw()
}
