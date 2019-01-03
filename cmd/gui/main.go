package main

import (
	"fmt"

	"github.com/ysh86/gui"
)

func main() {
	a := gui.NewApplication()
	if err := a.Init(); err != nil {
		panic(err)
	}
	defer a.Deinit()

	errc := a.Loop()
	select {
	case e := <-errc:
		if e != nil {
			panic(e)
		}
	}

	fmt.Println("Done")
}
