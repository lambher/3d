// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3 and OpenGL 2.1.
package main // import "github.com/go-gl/example/gl21-cube"

import (
	"go/build"
	_ "image/png"
	"log"
	"runtime"

	"github.com/lambher/3d/texture"

	"github.com/lambher/3d/entity/scene"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

//var (
//	texture   uint32
//	rotationX float32
//	rotationY float32
//)

const width, height = 800, 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(width, height, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	t := texture.NewTexture()
	t.Load()
	defer t.Delete()
	//texture = newTexture("./assets/grass_side.png")
	//defer gl.DeleteTextures(1, &texture)

	basicScene := scene.NewBasic(t, width, height)

	basicScene.Setup()
	for !window.ShouldClose() {
		basicScene.Draw()
		basicScene.Update()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// Set the working directory to the root of Go package, so that its assets can be accessed.
//func init() {
//	dir, err := importPathToDir("./assets")
//	if err != nil {
//		log.Fatalln("Unable to find Go package in your GOPATH, it's needed to load assets:", err)
//	}
//	err = os.Chdir(dir)
//	if err != nil {
//		log.Panicln("os.Chdir:", err)
//	}
//}

// importPathToDir resolves the absolute path from importPath.
// There doesn't need to be a valid Go package inside that import path,
// but the directory must exist.
func importPathToDir(importPath string) (string, error) {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return p.Dir, nil
}
