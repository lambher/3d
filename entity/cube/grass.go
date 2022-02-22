package cube

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lambher/3d/texture"
)

type Grass struct {
	*Cube
}

func NewGrass(t *texture.Texture, p mgl32.Vec3) *Grass {
	return &Grass{
		Cube: &Cube{
			textureSide:   t.Bucket[texture.GrassSide],
			textureTop:    t.Bucket[texture.GrassTop],
			textureBottom: t.Bucket[texture.GrassBottom],
			Position:      p,
		},
	}
}

func (g *Grass) Update() {
	//g.rotationX += 0.1
	g.rotationY -= 0.5
}

func (g Grass) Draw() {
	g.Cube.Draw()
}
