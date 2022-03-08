// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3 and OpenGL 2.1.
package main // import "github.com/go-gl/example/gl21-cube"

import (
	"fmt"
	_ "image/png"
	"log"
	"os"
	"runtime"

	"github.com/lambher/3d/assets/maps"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lambher/3d/context"
	"github.com/lambher/3d/controller"
	"github.com/lambher/3d/entity/scene"
)

const width, height = 1600, 900

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "seed" {
			err := maps.Seed()
			if err != nil {
				log.Fatalln("failed to seed:", err)
			}
			return
		}
	}

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
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

	ctx := context.NewContext()
	ctx.Texture.Load()
	defer ctx.Texture.Delete()

	basicScene, err := scene.NewBasic(ctx, os.Args[1])
	if err != nil {
		fmt.Println("cannot load scene", err)
		return
	}

	basicScene.Setup()
	for !window.ShouldClose() {
		controller.ListenInputs(window, ctx.Input)
		basicScene.Draw()
		basicScene.Update()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
