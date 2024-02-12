package main

import (
	"multiwinglscene/pkg/logger"
	"multiwinglscene/pkg/window"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	rotationX float32
	rotationY float32
)

const (
	width  = 800
	height = 600
	title  = "Go OpenGL Window"
)

func init() {
	// This is required to make sure GLFW and OpenGL calls run on the main thread.
	runtime.LockOSThread()
}

func main() {

	if err := glfw.Init(); err != nil {
		logger.LogError(err)
		return
		// log.Fatal("glfw.Init failed:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	var (
		windowsCount = 5
		windows      = make([]*window.Window, 0, windowsCount)
	)
	for i := 0; i < windowsCount; i++ {

		window, err := window.NewWindow(width, height, title, func() {
			gl.ClearColor(0, 0, 0, 0)
			gl.ClearDepth(1)
			gl.DepthFunc(gl.LEQUAL)

			gl.MatrixMode(gl.PROJECTION)
			gl.LoadIdentity()
			f := float64(width)/height - 1
			gl.Frustum(-1-f, 1+f, -1, 1, 1.0, 10.0)
			gl.MatrixMode(gl.MODELVIEW)
			gl.LoadIdentity()
		})
		if err != nil {
			logger.LogError(err)
			return
			// log.Fatal("NewWindow failed:", err)
		}
		windows = append(windows, window)
	}

	for {
		closedWindows := 0
		for i, w := range windows {
			if w.Closed() {
				closedWindows++
				continue
			}

			w.MakeContextCurrent()

			// Render OpenGL content here
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			drawCube(0, 0, 0.2*float32(i), 0.3, 0.3, 1.0)

			w.SwapBuffers()

			glfw.PollEvents()
		}
		if closedWindows == windowsCount {
			break
		}
		rotationX += 0.5
		rotationY += 0.5

	}
}

func drawCube(positionX, positionY, r, g, b, a float32) {

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(positionX, positionY, -3.0)
	gl.Rotatef(rotationX, 1, 0, 0)
	gl.Rotatef(rotationY, 0, 1, 0)

	gl.Color4f(r, g, b, a)

	gl.Begin(gl.LINE_LOOP)

	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, 1)

	gl.End()
	gl.Begin(gl.LINE_LOOP)

	gl.Normal3f(0, 0, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, -1)

	gl.End()
	gl.Begin(gl.LINE_LOOP)

	gl.Normal3f(0, 1, 0)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, -1)

	gl.End()
	gl.Begin(gl.LINE_LOOP)

	gl.Normal3f(0, -1, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, 1)

	gl.End()
	gl.Begin(gl.LINE_LOOP)

	gl.Normal3f(1, 0, 0)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, -1, -1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, -1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(1, -1, 1)

	gl.End()
	gl.Begin(gl.LINE_LOOP)

	gl.Normal3f(-1, 0, 0)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, -1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(-1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, -1)

	gl.End()
}
