package models

type CubeType string

const (
	Grass  CubeType = "grass"
	Grass2 CubeType = "grass2"
)

type Cube struct {
	Position Vector3  `json:"position"`
	Type     CubeType `json:"type"`
}
