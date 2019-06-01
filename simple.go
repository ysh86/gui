// +build !windows

package gui

import (
	"log"
	"os"
	"runtime"
)

type application struct {
	logger *log.Logger
	// ...
}

// NewApplication creates a new GUI application.
func NewApplication() Application {
	return &application{}
}

func (a *application) EnableLog() error {
	a.logger = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	if a.logger != nil {
		a.logger.Print("start logging")
	}
	return nil
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
