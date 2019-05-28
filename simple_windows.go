package gui

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"unsafe"

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

	hwnds     []windows.Handle
	renderers []*Renderer
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
	className := "GO GUI: simple window app"
	classNameUTF16, err := windows.UTF16PtrFromString(className)
	if err != nil {
		return fmt.Errorf("UTF16PtrFromString %s: %v", className, err)
	}
	icon, err := LoadIcon(0, MAKEINTRESOURCE(IDI_APPLICATION))
	if err != nil {
		return fmt.Errorf("LoadIcon: %v", err)
	}
	cursor, err := LoadCursor(0, MAKEINTRESOURCE(IDC_ARROW))
	if err != nil {
		return fmt.Errorf("LoadCursor: %v", err)
	}
	wndClass := &WndClassEx{
		Size:       0,
		Style:      CS_HREDRAW | CS_VREDRAW,
		WndProc:    windows.NewCallback(a.windowProc),
		ClsExtra:   0,
		WndExtra:   0,
		Instance:   a.instance,
		Icon:       icon,
		Cursor:     cursor,
		Background: windows.Handle(COLOR_WINDOW + 1),
		MenuName:   nil,
		ClassName:  classNameUTF16,
		IconSm:     0,
	}
	wndClass.Size = uint32(unsafe.Sizeof(*wndClass))
	atom, err := RegisterClassEx(wndClass)
	if err != nil {
		return fmt.Errorf("RegisterClassEx %v: %v", wndClass, err)
	}
	a.atom = atom

	return nil
}

func (a *application) Deinit() error {
	if a != nil {
		/*
			for _, r := range a.renderers {
				// nothing to do
			}
		*/
	}

	return nil
}

func (a *application) Loop(windowName string, renderer Renderer) <-chan error {
	errc := make(chan error, 1)

	go func() {
		// lock thread for the GetMessage() API
		runtime.LockOSThread()

		w, err := a.appendWindow(windowName, renderer)
		if err != nil {
			errc <- err
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

			// debug
			fmt.Fprintf(os.Stderr, "GetMessage: %p, 0x%08x, %p\n", unsafe.Pointer(w), msg.message, unsafe.Pointer(msg.wParam))
			if result == 0 {
				// WM_QUIT (wParam is ExitCode)
				if msg.wParam == 0 {
					errc <- nil
				} else {
					errc <- fmt.Errorf("GetMessage: %p, WM_QUIT, %d", unsafe.Pointer(w), msg.wParam)
				}
				break
			}

			TranslateMessage(&msg)
			DispatchMessage(&msg)
		}
	}()

	return errc
}

func (a *application) appendWindow(name string, renderer Renderer) (windows.Handle, error) {
	nameUTF16, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return 0, fmt.Errorf("UTF16PtrFromString %s: %v", name, err)
	}

	w, err := CreateWindowEx(
		0,
		(*uint16)(unsafe.Pointer(uintptr(a.atom))),
		nameUTF16,
		WS_OVERLAPPEDWINDOW,
		CW_USEDEFAULT, CW_USEDEFAULT,
		CW_USEDEFAULT, CW_USEDEFAULT, // W x H
		0,
		0,
		a.instance,
		uintptr(unsafe.Pointer(&renderer)),
	)
	if err != nil {
		return 0, fmt.Errorf("CreateWindowEx: %v", err)
	}
	a.hwnds = append(a.hwnds, w)
	a.renderers = append(a.renderers, &renderer)

	_ = ShowWindow(w, a.cmdShow) // ignore return value
	_ = UpdateWindow(w)          // ignore return value

	return w, err
}

// TODO: debug flag 付けよう
func (a *application) windowProc(window windows.Handle, message uint32, wParam uintptr, lParam uintptr) uintptr {
	// debug
	fmt.Fprintf(os.Stderr, "windowProc: %p, 0x%08x\n", unsafe.Pointer(window), message)

	// save renderer as user data
	if message == WM_CREATE {
		cs := (*CreateStruct)(unsafe.Pointer(lParam))
		ptr := cs.CreateParams

		// check ptr is a valid Renderer
		if ptr != 0 {
			renderer := (*Renderer)(unsafe.Pointer(ptr))
			raw := reflect.ValueOf(*renderer)
			if !raw.IsValid() || raw.Kind() != reflect.Ptr || raw.IsNil() {
				ptr = 0
			}
		}

		SetWindowLongPtr(
			window,
			GWLP_USERDATA,
			ptr,
		)

		return 1
	}

	// use user data as renderer
	ptr, err := GetWindowLongPtr(
		window,
		GWLP_USERDATA,
	)
	if err != nil {
		ptr = 0
	}
	renderer := (*Renderer)(unsafe.Pointer(ptr))

	switch message {
	case WM_PAINT:
		if renderer != nil {
			// draw something...
		}
		return 0
	case WM_DESTROY:
		PostQuitMessage(0)
		return 1
	}

	r, _ := DefWindowProc(window, message, wParam, lParam)
	return r
}
