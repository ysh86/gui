package gui

// Application is the GUI application.
type Application interface {
	Init() error
	Deinit()
	EnableLog() error
	Loop(windowName string, width int32, height int32, renderer Renderer) <-chan error
}

// Renderer is a renderer for drawing window contents.
type Renderer interface {
	Init() error
	Deinit()
	Dpi() (float32, float32)
	Update(width, height uint32) error
	Draw(nativeWindow uintptr) error
}

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -systemdll -output zgui_windows.go gui_windows.go
