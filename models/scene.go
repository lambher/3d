package models

type Scene struct {
	Light  Light  `json:"light"`
	Camera Camera `json:"camera"`
	Cubes  []Cube `json:"cubes"`
}
