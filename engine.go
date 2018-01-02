package plane

import (
	"image/color"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

// Engine contains the internal state of the engine.
type Engine struct {
	window *pixelgl.Window
	root   Object
	vals   map[string]interface{}
}

// Main runs the engine's main loop. It calls init to initialize the
// engine, then enters the loop.
//
// Note that the init function *must* call the Engine instance's
// CreateWindow method. Failure to do so will result in a panic after
// the init function returns.
func Main(init func(e *Engine)) {
	pixelgl.Run(func() {
		e := Engine{
			vals: make(map[string]interface{}),
		}

		init(&e)

		if e.window == nil {
			panic("init function did not create window")
		}
		defer e.window.Destroy()

		fps := time.NewTicker(time.Second / 60)
		defer fps.Stop()

		for !e.window.Closed() {
			<-fps.C

			root := e.root
			root.Update(&e)
			e.window.Clear(color.Black)
			root.Draw(&e)
			e.window.Update()
		}
	})
}

// CreateWindow creates the main window for the engine. It must be
// called exactly once during the init function that is passed to
// Main.
func (e *Engine) CreateWindow(cfg pixelgl.WindowConfig) error {
	if e.window != nil {
		panic("Window already created")
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return err
	}
	e.window = win
	return nil
}

// Window returns the main window. Note that this will return nil
// until if called before CreateWindow.
func (e *Engine) Window() *pixelgl.Window {
	return e.window
}

func (e *Engine) SetRoot(obj Object) {
	e.root = obj
}

// Get gets the global value associated with the given key.
func (e *Engine) Get(key string) interface{} {
	return e.vals[key]
}

// Set associates a global value with the given key.
func (e *Engine) Set(key string, val interface{}) {
	e.vals[key] = val
}
