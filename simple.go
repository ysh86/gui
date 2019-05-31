// +build !windows

package gui

import (
	"runtime"
)

type application struct {
	// ...
}

// NewApplication creates a new GUI application.
func NewApplication() Application {
	return &application{}
}

func (a *application) Init() error {
	return nil
}

func (a *application) Deinit() {
	if a != nil {
		// nothing to do
	}
}

func (a *application) Loop(windowName string, width int32, height int32, renderer Renderer) <-chan error {
	errc := make(chan error, 1)

	go func() {
		// lock thread for message handling
		runtime.LockOSThread()

		// create a window
		// ...

		// message loop
		for {
			// drawing...

			// quit
			errc <- nil
			break
		}
	}()

	return errc
}
