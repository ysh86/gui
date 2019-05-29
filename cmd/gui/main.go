package main

import (
	"fmt"
	"os"

	"github.com/ysh86/gui"
)

func main() {
	app := gui.NewApplication()

	if err := app.Init(); err != nil {
		panic(err)
	}
	defer app.Deinit()

	windowName := "single window"
	errc := app.Loop(windowName, 640, 480, nil)
	select {
	case e := <-errc:
		if e != nil {
			panic(e)
		}
	}

	fmt.Fprintln(os.Stderr, "Done")
}
