package gui

// Application is the GUI application.
type Application interface {
	Init() error
	Deinit() error
	Loop(windowName string, renderer Renderer) <-chan error
}

// Renderer is a renderer for drawing window contents.
type Renderer interface {
	Init() error
	Deinit() error
	Update() error
	Draw() error
}

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -systemdll -output zgui_windows.go gui_windows.go
