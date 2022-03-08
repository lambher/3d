package controller

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lambher/3d/input"
)

var keyMap = make(map[input.Key]glfw.Key)

func init() {
	keyMap[input.MoveForward] = glfw.KeyW
	keyMap[input.MoveBackward] = glfw.KeyS
	keyMap[input.MoveLeft] = glfw.KeyA
	keyMap[input.MoveRight] = glfw.KeyD
	keyMap[input.MoveUp] = glfw.KeySpace
	keyMap[input.MoveDown] = glfw.KeyLeftShift

	keyMap[input.LookLeft] = glfw.KeyLeft
	keyMap[input.LookRight] = glfw.KeyRight
	keyMap[input.LookUp] = glfw.KeyUp
	keyMap[input.LookDown] = glfw.KeyDown
}

func ListenInputs(window *glfw.Window, i *input.Input) {
	for key, value := range keyMap {
		handleInput(window, i, key, value)
	}
}

func handleInput(window *glfw.Window, i *input.Input, key input.Key, value glfw.Key) {
	if window.GetKey(value) == glfw.Press {
		i.Keys[key] = true
	}
	if window.GetKey(value) == glfw.Release {
		i.Keys[key] = false
	}

}
