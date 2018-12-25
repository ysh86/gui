package gui

type Application interface {
	Init() error
	Loop() <-chan error
}

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -systemdll -output zgui_windows.go gui_windows.go
