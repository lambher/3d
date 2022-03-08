package models

type Camera struct {
	Position Vector3 `json:"position"`
	View     Vector2 `json:"view"`
}
