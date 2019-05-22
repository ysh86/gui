// Windows gui system calls.

package gui

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// objbase.h
const (
	COINIT_APARTMENTTHREADED = 2
	COINIT_MULTITHREADED     = 0
	COINIT_DISABLE_OLE1DDE   = 4
	COINIT_SPEED_OVER_MEMORY = 8
)

// winuser.h
const (
	// ShowWindow() codes
	SW_HIDE            = 0
	SW_SHOWNORMAL      = 1
	SW_NORMAL          = SW_SHOWNORMAL
	SW_SHOWMINIMIZED   = 2
	SW_SHOWMAXIMIZED   = 3
	SW_MAXIMIZE        = SW_SHOWMAXIMIZED
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
	SW_MAX             = 11
)
const (
	// Class styles
	CS_VREDRAW         = 0x00000001
	CS_HREDRAW         = 0x00000002
	CS_DBLCLKS         = 0x00000008
	CS_OWNDC           = 0x00000020
	CS_CLASSDC         = 0x00000040
	CS_PARENTDC        = 0x00000080
	CS_NOCLOSE         = 0x00000200
	CS_SAVEBITS        = 0x00000800
	CS_BYTEALIGNCLIENT = 0x00001000
	CS_BYTEALIGNWINDOW = 0x00002000
	CS_GLOBALCLASS     = 0x00004000
	CS_IME             = 0x00010000
	CS_DROPSHADOW      = 0x00020000
)
const (
	COLOR_SCROLLBAR           = 0
	COLOR_BACKGROUND          = 1
	COLOR_ACTIVECAPTION       = 2
	COLOR_INACTIVECAPTION     = 3
	COLOR_MENU                = 4
	COLOR_WINDOW              = 5
	COLOR_WINDOWFRAME         = 6
	COLOR_MENUTEXT            = 7
	COLOR_WINDOWTEXT          = 8
	COLOR_CAPTIONTEXT         = 9
	COLOR_ACTIVEBORDER        = 10
	COLOR_INACTIVEBORDER      = 11
	COLOR_APPWORKSPACE        = 12
	COLOR_HIGHLIGHT           = 13
	COLOR_HIGHLIGHTTEXT       = 14
	COLOR_BTNFACE             = 15
	COLOR_BTNSHADOW           = 16
	COLOR_GRAYTEXT            = 17
	COLOR_BTNTEXT             = 18
	COLOR_INACTIVECAPTIONTEXT = 19
	COLOR_BTNHIGHLIGHT        = 20

	COLOR_3DDKSHADOW              = 21
	COLOR_3DLIGHT                 = 22
	COLOR_INFOTEXT                = 23
	COLOR_INFOBK                  = 24
	COLOR_HOTLIGHT                = 26
	COLOR_GRADIENTACTIVECAPTION   = 27
	COLOR_GRADIENTINACTIVECAPTION = 28
	COLOR_MENUHILIGHT             = 29
	COLOR_MENUBAR                 = 30

	COLOR_DESKTOP     = COLOR_BACKGROUND
	COLOR_3DFACE      = COLOR_BTNFACE
	COLOR_3DSHADOW    = COLOR_BTNSHADOW
	COLOR_3DHIGHLIGHT = COLOR_BTNHIGHLIGHT
	COLOR_3DHILIGHT   = COLOR_BTNHIGHLIGHT
	COLOR_BTNHILIGHT  = COLOR_BTNHIGHLIGHT
)
const (
	// Window styles
	WS_OVERLAPPED   = 0x00000000
	WS_POPUP        = 0x80000000
	WS_CHILD        = 0x40000000
	WS_MINIMIZE     = 0x20000000
	WS_VISIBLE      = 0x10000000
	WS_DISABLED     = 0x08000000
	WS_CLIPSIBLINGS = 0x04000000
	WS_CLIPCHILDREN = 0x02000000
	WS_MAXIMIZE     = 0x01000000
	WS_CAPTION      = (WS_BORDER | WS_DLGFRAME)
	WS_BORDER       = 0x00800000
	WS_DLGFRAME     = 0x00400000
	WS_VSCROLL      = 0x00200000
	WS_HSCROLL      = 0x00100000
	WS_SYSMENU      = 0x00080000
	WS_THICKFRAME   = 0x00040000
	WS_GROUP        = 0x00020000
	WS_TABSTOP      = 0x00010000
	WS_MINIMIZEBOX  = 0x00020000
	WS_MAXIMIZEBOX  = 0x00010000
	WS_TILED        = WS_OVERLAPPED
	WS_ICONIC       = WS_MINIMIZE
	WS_SIZEBOX      = WS_THICKFRAME
	WS_TILEDWINDOW  = WS_OVERLAPPEDWINDOW

	WS_OVERLAPPEDWINDOW = (WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX)
	WS_POPUPWINDOW      = (WS_POPUP | WS_BORDER | WS_SYSMENU)
	WS_CHILDWINDOW      = WS_CHILD
)
const (
	// CreateWindow() coordinates
	CW_USEDEFAULT = -2147483648 // = 0x80000000
)
const (
	// Messages
	WM_DESTROY       = 0x0002
	WM_SIZE          = 0x0005
	WM_PAINT         = 0x000F
	WM_DISPLAYCHANGE = 0x007E
)
const (
	// Icons
	IDI_APPLICATION = 32512
	IDI_HAND        = 32513
	IDI_QUESTION    = 32514
	IDI_EXCLAMATION = 32515
	IDI_ASTERISK    = 32516
	IDI_WINLOGO     = 32517
	IDI_SHIELD      = 32518

	IDI_WARNING     = IDI_EXCLAMATION
	IDI_ERROR       = IDI_HAND
	IDI_INFORMATION = IDI_ASTERISK
)
const (
	// Cursors
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZE        = 32640
	IDC_ICON        = 32641
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
)

// macros
func MAKEINTRESOURCE(value uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(value)))
}
func LOWORD(value uintptr) uint16 {
	return uint16(value & 0xFFFF)
}
func HIWORD(value uintptr) uint16 {
	return uint16((value >> 16) & 0xFFFF)
}

// WndClassEx is a struct for RegisterClassEx().
type WndClassEx struct {
	Size uint32

	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   windows.Handle
	Icon       windows.Handle
	Cursor     windows.Handle
	Background windows.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     windows.Handle
}

// Atom is a returned value from RegisterClassEx()
type Atom uint16

// Point holds x and y.
type Point struct {
	X int32
	Y int32
}

// Rect holds left, top, right and bottom.
type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// Msg is a message struct for the message loop.
type Msg struct {
	hwnd    windows.Handle
	message uint32
	wparam  uintptr
	lparam  uintptr
	time    uint32
	pt      Point
}

// windows api calls

//sys	GetModuleHandle(modulename *uint16) (module windows.Handle, err error) = GetModuleHandleW
//sys	CoInitializeEx(reserved uintptr, coInit uint32) (err error) [failretval!=0] = ole32.CoInitializeEx
//sys	CoUninitialize() = ole32.CoUninitialize
//sys	MessageBoxEx(window windows.Handle, text *uint16, caption *uint16, style uint32, languageID uint16) (id int32, err error) = user32.MessageBoxExW
//sys   LoadIcon(instance windows.Handle, iconName *uint16) (icon windows.Handle, err error) [failretval==0] = user32.LoadIconW
//sys   LoadCursor(instance windows.Handle, cursorName *uint16) (cursor windows.Handle, err error) [failretval==0] = user32.LoadCursorW
//sys	RegisterClassEx(class *WndClassEx) (atom Atom, err error) = user32.RegisterClassExW
//sys	CreateWindowEx(exStyle uint32, classname *uint16, windowname *uint16, style uint32, x int32, y int32, width int32, height int32, parent windows.Handle, menu windows.Handle, instance windows.Handle, lparam uintptr) (window windows.Handle, err error) = user32.CreateWindowExW
//sys	ShowWindow(window windows.Handle, command int32) (err error) [failretval!=0] = user32.ShowWindow
//sys	UpdateWindow(window windows.Handle) (err error) = user32.UpdateWindow
//sys	DefWindowProc(window windows.Handle, message uint32, wparam uintptr, lparam uintptr) (result uintptr, err error) = user32.DefWindowProcW
//sys	GetMessage(message *Msg, window windows.Handle, messageFilterMin uint32, messageFilterMax uint32) (result int32, err error) [failretval==-1] = user32.GetMessageW
//sys	TranslateMessage(message *Msg) (err error) = user32.TranslateMessage
//sys	DispatchMessage(message *Msg) (result uintptr, err error) = user32.DispatchMessageW
//sys	PostQuitMessage(exitCode int32) = user32.PostQuitMessage
//sys	GetClientRect(window windows.Handle, rect *Rect) (err error) [failretval==0] = user32.GetClientRect
//sys	ValidateRect(window windows.Handle, rect *Rect) (err error) [failretval==0] = user32.ValidateRect
//sys	InvalidateRect(window windows.Handle, rect *Rect, erase bool) (err error) [failretval==0] = user32.InvalidateRect
