package pixel

import (
	"roguespace/plane"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type backendImpl struct {
}

func Backend() plane.Backend {
	return backendImpl{}
}

func (b backendImpl) CreateWindow(cfg plane.WindowConfig) (plane.Window, error) {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  cfg.Title,
		Bounds: pixel.R(0, 0, cfg.Size.X, cfg.Size.Y),
	})
	if err != nil {
		return nil, err
	}
	return windowImpl{win}, nil
}

type windowImpl struct {
	*pixelgl.Window
}

func (w windowImpl) Pressed(b plane.Button) bool {
	return w.Window.Pressed(keyToPixel[b])
}

func (w windowImpl) JustPressed(b plane.Button) bool {
	return w.Window.JustPressed(keyToPixel[b])
}

func (w windowImpl) JustReleased(b plane.Button) bool {
	return w.Window.JustReleased(keyToPixel[b])
}

func (w windowImpl) MousePos() plane.Vector {
	return toVector(w.Window.MousePosition())
}

func (w windowImpl) MouseScroll() plane.Vector {
	return toVector(w.Window.MouseScroll())
}

func toVector(v pixel.Vec) plane.Vector {
	return plane.Vector(v)
}
