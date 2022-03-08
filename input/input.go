package input

type Key string

const (
	MoveForward  Key = "move_forward"
	MoveBackward Key = "move_backward"
	MoveLeft     Key = "move_left"
	MoveRight    Key = "move_right"
	MoveUp       Key = "move_up"
	MoveDown     Key = "move_down"

	LookLeft  Key = "look_left"
	LookRight Key = "look_right"
	LookUp    Key = "look_up"
	LookDown  Key = "look_down"
)

type Input struct {
	Keys map[Key]bool
}

func NewInput() *Input {
	return &Input{
		Keys: make(map[Key]bool),
	}
}
