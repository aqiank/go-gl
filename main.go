// package main

import (
	"github.com/jackyb/go-sdl2/sdl"
	"./gl"
)

func main() {
	window := sdl.CreateWindow("go-gl", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			1024, 768, sdl.WINDOW_OPENGL);
	glcontext := sdl.GL_CreateContext(window)
	gl.Begin(gl.QUADS)
	gl.Color3f(1, 0, 0)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(1, 0)
	gl.Vertex2f(1, 1)
	gl.Vertex2f(0, 1)
	gl.End()
	sdl.GL_SwapWindow(window)
	sdl.Delay(1000)
	sdl.GL_DeleteContext(glcontext)
	window.Destroy()
}
