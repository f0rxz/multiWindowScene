package window

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	destroyed bool
	*glfw.Window
}

func NewWindow(width, height int, title string, setupScene func()) (*Window, error) {
	glfwWindow, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}

	glfwWindow.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		return nil, err
	}

	setupScene()

	return &Window{
		Window:    glfwWindow,
		destroyed: false,
	}, nil
}

func (w *Window) Closed() bool {
	if w.destroyed {
		return true
	}

	if w.ShouldClose() {
		w.Destroy()
		w.destroyed = true
		return true
	}

	return false
}
