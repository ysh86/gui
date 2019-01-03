package gui

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ysh86/d2d1"
	"golang.org/x/sys/windows"
)

func init() {
	captionUTF16, _ := windows.UTF16PtrFromString("Hello")
	textUTF16, _ := windows.UTF16PtrFromString("from init(): Hello 世界")
	MessageBoxEx(0, textUTF16, captionUTF16, 0, 0)
}

type application struct {
	instance windows.Handle
	cmdLine  string
	cmdShow  int32
	atom     Atom
	factory  *d2d1.Factory

	wnds []windows.Handle
}

// NewApplication creates a new GUI application.
func NewApplication() Application {
	return &application{}
}

func (a *application) Init() error {
	// dummy _tWinMain()
	i, err := GetModuleHandle(nil)
	if err != nil {
		return fmt.Errorf("GetModuleHandle: %v", err)
	}
	a.instance = i
	a.cmdLine = ""
	a.cmdShow = SW_SHOWNORMAL

	// register a window class
	className := "This is a simple window app."
	classNameUTF16, err := windows.UTF16PtrFromString(className)
	if err != nil {
		return fmt.Errorf("UTF16PtrFromString %s: %v", className, err)
	}
	wndClass := &WndClassEx{
		size:       0,
		style:      CS_HREDRAW | CS_VREDRAW,
		wndProc:    windows.NewCallback(windowProc),
		clsExtra:   0,
		wndExtra:   0,
		instance:   a.instance,
		icon:       0,
		cursor:     0,
		background: windows.Handle(COLOR_WINDOW + 1),
		menuName:   nil,
		className:  classNameUTF16,
		iconSm:     0,
	}
	wndClass.size = uint32(unsafe.Sizeof(*wndClass))
	atom, err := RegisterClassEx(wndClass)
	if err != nil {
		return fmt.Errorf("RegisterClassEx %v: %v", wndClass, err)
	}
	a.atom = atom

	// D2D1
	factory, err := d2d1.CreateFactory(d2d1.FACTORY_TYPE_SINGLE_THREADED, nil)
	if err != nil {
		return err
	}
	a.factory = factory

	return nil
}

func (a *application) Deinit() error {
	if a != nil {
		if a.factory != nil {
			a.factory.Release()
		}
	}

	return nil
}

func (a *application) Loop() <-chan error {
	errc := make(chan error, 1)

	go func() {
		// lock thread for the GetMessage() API
		runtime.LockOSThread()

		// create a window
		windowName := "single window"
		windowNameUTF16, err := windows.UTF16PtrFromString(windowName)
		if err != nil {
			errc <- fmt.Errorf("UTF16PtrFromString %s: %v", windowName, err)
			return
		}
		w, err := CreateWindowEx(
			0,
			(*uint16)(unsafe.Pointer(uintptr(a.atom))),
			windowNameUTF16,
			WS_OVERLAPPEDWINDOW,
			CW_USEDEFAULT, CW_USEDEFAULT,
			CW_USEDEFAULT, CW_USEDEFAULT, // W x H
			0,
			0,
			0,
			0,
		)
		if err != nil {
			errc <- fmt.Errorf("CreateWindowEx: %v", err)
			return
		}
		a.wnds = append(a.wnds, w)

		_ = ShowWindow(w, a.cmdShow) // ignore return value
		err = UpdateWindow(w)
		if err != nil {
			errc <- fmt.Errorf("UpdateWindow %p: %v", unsafe.Pointer(w), err)
			return
		}

		// message loop
		var msg Msg
		for {
			result, err := GetMessage(&msg, 0, 0, 0)
			if err != nil {
				errc <- fmt.Errorf("GetMessage: %p, %v", unsafe.Pointer(w), err)
				return
			}

			fmt.Printf("GetMessage: %p, 0x%08x, %p\n", unsafe.Pointer(w), msg.message, unsafe.Pointer(msg.wparam))
			if result == 0 {
				// WM_QUIT (wparam is ExitCode)
				if msg.wparam == 0 {
					errc <- nil
				} else {
					errc <- fmt.Errorf("GetMessage: %p, WM_QUIT, %d", unsafe.Pointer(w), msg.wparam)
				}
				break
			}

			TranslateMessage(&msg)
			DispatchMessage(&msg)
		}
	}()

	return errc
}

func windowProc(window windows.Handle, message uint32, wparam uintptr, lparam uintptr) uintptr {
	fmt.Printf("windowProc: %p, 0x%08x\n", unsafe.Pointer(window), message)

	var result uintptr

	switch message {
	case WM_DESTROY:
		PostQuitMessage(0)
	default:
		r, _ := DefWindowProc(window, message, wparam, lparam)
		result = r
	}

	return result
}
