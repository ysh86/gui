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

func (a *application) Deinit() error {
	return nil
}

func (a *application) Loop() <-chan error {
	errc := make(chan error, 1)

	go func() {
		// lock thread for message handling
		runtime.LockOSThread()

		// create a window
		// ...

		// message loop
		for {
			// ...

			// quit
			errc <- nil
			break
		}
	}()

	return errc
}
