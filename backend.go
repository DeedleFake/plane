package plane

import "image/color"

type Backend interface {
	CreateWindow(WindowConfig) (Window, error)
}

type Window interface {
	Target

	Pressed(Button) bool
	JustPressed(Button) bool
	JustReleased(Button) bool
	Typed() string
	MousePos() Vector
	MouseScroll() Vector

	Clear(color.Color)
	Update()
}

type WindowConfig struct {
	Title string
	Size  Vector
}

type Target interface{}
