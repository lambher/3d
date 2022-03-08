package entity

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lambher/3d/models"
)

type Entity interface {
	Draw()
	Update()
}

func Vec3(v models.Vector3) mgl32.Vec3 {
	return mgl32.Vec3{v.X, v.Y, v.Z}
}
