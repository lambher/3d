package context

import (
	"github.com/lambher/3d/input"
	"github.com/lambher/3d/texture"
)

type Context struct {
	Texture *texture.Texture
	Input   *input.Input
}

func NewContext() *Context {
	return &Context{
		Texture: texture.NewTexture(),
		Input:   input.NewInput(),
	}
}

func (c Context) GetTexture(name texture.Name) uint32 {
	return c.Texture.Bucket[name]
}

func (c Context) GetInputValue(key input.Key) bool {
	return c.Input.Keys[key]
}
