// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package sys

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	moduser32   = windows.NewLazySystemDLL("user32.dll")
	modgdi32    = windows.NewLazySystemDLL("gdi32.dll")
	modshell32  = windows.NewLazySystemDLL("shell32.dll")

	procGetModuleHandleW       = modkernel32.NewProc("GetModuleHandleW")
	procVerifyVersionInfoW     = modkernel32.NewProc("VerifyVersionInfoW")
	procVerSetConditionMask    = modkernel32.NewProc("VerSetConditionMask")
	procAppendMenuW            = moduser32.NewProc("AppendMenuW")
	procCreateIconIndirect     = moduser32.NewProc("CreateIconIndirect")
	procCreatePopupMenu        = moduser32.NewProc("CreatePopupMenu")
	procCreateWindowExW        = moduser32.NewProc("CreateWindowExW")
	procDefWindowProcW         = moduser32.NewProc("DefWindowProcW")
	procDestroyIcon            = moduser32.NewProc("DestroyIcon")
	procDestroyMenu            = moduser32.NewProc("DestroyMenu")
	procDestroyWindow          = moduser32.NewProc("DestroyWindow")
	procDispatchMessageW       = moduser32.NewProc("DispatchMessageW")
	procGetCursorPos           = moduser32.NewProc("GetCursorPos")
	procGetDC                  = moduser32.NewProc("GetDC")
	procGetMessageW            = moduser32.NewProc("GetMessageW")
	procGetWindowLongW         = moduser32.NewProc("GetWindowLongW")
	procGetWindowLongPtrW      = moduser32.NewProc("GetWindowLongPtrW")
	procLoadImageW             = moduser32.NewProc("LoadImageW")
	procPostMessageW           = moduser32.NewProc("PostMessageW")
	procPostQuitMessage        = moduser32.NewProc("PostQuitMessage")
	procRegisterClassExW       = moduser32.NewProc("RegisterClassExW")
	procRegisterWindowMessageW = moduser32.NewProc("RegisterWindowMessageW")
	procReleaseDC              = moduser32.NewProc("ReleaseDC")
	procSetForegroundWindow    = moduser32.NewProc("SetForegroundWindow")
	procSetWindowLongW         = moduser32.NewProc("SetWindowLongW")
	procSetWindowLongPtrW      = moduser32.NewProc("SetWindowLongPtrW")
	procTrackPopupMenu         = moduser32.NewProc("TrackPopupMenu")
	procTranslateMessage       = moduser32.NewProc("TranslateMessage")
	procCreateCompatibleBitmap = modgdi32.NewProc("CreateCompatibleBitmap")
	procCreateCompatibleDC     = modgdi32.NewProc("CreateCompatibleDC")
	procDeleteDC               = modgdi32.NewProc("DeleteDC")
	procDeleteObject           = modgdi32.NewProc("DeleteObject")
	procSelectObject           = modgdi32.NewProc("SelectObject")
	procSetPixel               = modgdi32.NewProc("SetPixel")
	procShell_NotifyIconW      = modshell32.NewProc("Shell_NotifyIconW")
)

func GetModuleHandle(name *uint16) (h windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetModuleHandleW.Addr(), 1, uintptr(unsafe.Pointer(name)), 0, 0)
	h = windows.Handle(r0)
	if h == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func VerifyVersionInfo(vi *OSVersionInfoEx, typeMask uint32, conditionMask uint64) (ok bool) {
	r0, _, _ := syscall.Syscall(procVerifyVersionInfoW.Addr(), 3, uintptr(unsafe.Pointer(vi)), uintptr(typeMask), uintptr(conditionMask))
	ok = r0 != 0
	return
}

func VerSetConditionMask(lConditionMask uint64, typeBitMask uint32, conditionMask uint8) (mask uint64) {
	r0, _, _ := syscall.Syscall(procVerSetConditionMask.Addr(), 3, uintptr(lConditionMask), uintptr(typeBitMask), uintptr(conditionMask))
	mask = uint64(r0)
	return
}

func AppendMenu(menu windows.Handle, flags uint32, id uintptr, item *uint16) (err error) {
	r1, _, e1 := syscall.Syscall6(procAppendMenuW.Addr(), 4, uintptr(menu), uintptr(flags), uintptr(id), uintptr(unsafe.Pointer(item)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateIconIndirect(ii *IconInfo) (icon windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateIconIndirect.Addr(), 1, uintptr(unsafe.Pointer(ii)), 0, 0)
	icon = windows.Handle(r0)
	if icon == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreatePopupMenu() (menu windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreatePopupMenu.Addr(), 0, 0, 0, 0)
	menu = windows.Handle(r0)
	if menu == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateWindowEx(exStyle uint32, className *uint16, windowName *uint16, style uint32, x int32, y int32, w int32, h int32, parent windows.Handle, menu windows.Handle, inst windows.Handle, param unsafe.Pointer) (wnd windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExW.Addr(), 12, uintptr(exStyle), uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(windowName)), uintptr(style), uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(parent), uintptr(menu), uintptr(inst), uintptr(param))
	wnd = windows.Handle(r0)
	if wnd == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DefWindowProc(wnd windows.Handle, msg uint32, wParam uintptr, lParam uintptr) (res uintptr) {
	r0, _, _ := syscall.Syscall6(procDefWindowProcW.Addr(), 4, uintptr(wnd), uintptr(msg), uintptr(wParam), uintptr(lParam), 0, 0)
	res = uintptr(r0)
	return
}

func DestroyIcon(icon windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyIcon.Addr(), 1, uintptr(icon), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DestroyMenu(menu windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyMenu.Addr(), 1, uintptr(menu), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DestroyWindow(wnd windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyWindow.Addr(), 1, uintptr(wnd), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DispatchMessage(msg *Msg) (res uintptr) {
	r0, _, _ := syscall.Syscall(procDispatchMessageW.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	res = uintptr(r0)
	return
}

func GetCursorPos(pt *Point) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCursorPos.Addr(), 1, uintptr(unsafe.Pointer(pt)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetDC(wnd windows.Handle) (dc windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetDC.Addr(), 1, uintptr(wnd), 0, 0)
	dc = windows.Handle(r0)
	if dc == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetMessage(msg *Msg, wnd windows.Handle, msgFilterMin uint32, msgFilterMax uint32) (ret int32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetMessageW.Addr(), 4, uintptr(unsafe.Pointer(msg)), uintptr(wnd), uintptr(msgFilterMin), uintptr(msgFilterMax), 0, 0)
	ret = int32(r0)
	if ret == -1 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getWindowLong(wnd windows.Handle, i int32) (ptr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowLongW.Addr(), 2, uintptr(wnd), uintptr(i), 0)
	ptr = uintptr(r0)
	if ptr == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getWindowLongPtr(wnd windows.Handle, i int32) (ptr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowLongPtrW.Addr(), 2, uintptr(wnd), uintptr(i), 0)
	ptr = uintptr(r0)
	if ptr == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func LoadImage(inst windows.Handle, name *uint16, typ uint32, cxDesired int32, cyDesired int32, load uint32) (h windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procLoadImageW.Addr(), 6, uintptr(inst), uintptr(unsafe.Pointer(name)), uintptr(typ), uintptr(cxDesired), uintptr(cyDesired), uintptr(load))
	h = windows.Handle(r0)
	if h == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func PostMessage(wnd windows.Handle, msg uint32, wParam uintptr, lParam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procPostMessageW.Addr(), 4, uintptr(wnd), uintptr(msg), uintptr(wParam), uintptr(lParam), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(procPostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
	return
}

func RegisterClassEx(wcx *WndClassEx) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassExW.Addr(), 1, uintptr(unsafe.Pointer(wcx)), 0, 0)
	atom = uint16(r0)
	if atom == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func RegisterWindowMessage(s *uint16) (msg uint32, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterWindowMessageW.Addr(), 1, uintptr(unsafe.Pointer(s)), 0, 0)
	msg = uint32(r0)
	if msg == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ReleaseDC(wnd windows.Handle, dc windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procReleaseDC.Addr(), 2, uintptr(wnd), uintptr(dc), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetForegroundWindow(wnd windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetForegroundWindow.Addr(), 1, uintptr(wnd), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setWindowLong(wnd windows.Handle, i int32, ptr unsafe.Pointer) (oldptr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procSetWindowLongW.Addr(), 3, uintptr(wnd), uintptr(i), uintptr(ptr))
	oldptr = uintptr(r0)
	if oldptr == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setWindowLongPtr(wnd windows.Handle, i int32, ptr unsafe.Pointer) (oldptr uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procSetWindowLongPtrW.Addr(), 3, uintptr(wnd), uintptr(i), uintptr(ptr))
	oldptr = uintptr(r0)
	if oldptr == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func TrackPopupMenu(menu windows.Handle, flags uint32, x int32, y int32, reserved int32, wnd windows.Handle) (ret int32, err error) {
	r0, _, e1 := syscall.Syscall6(procTrackPopupMenu.Addr(), 6, uintptr(menu), uintptr(flags), uintptr(x), uintptr(y), uintptr(reserved), uintptr(wnd))
	ret = int32(r0)
	if ret == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func TranslateMessage(msg *Msg) (err error) {
	r1, _, e1 := syscall.Syscall(procTranslateMessage.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateCompatibleBitmap(dc windows.Handle, w int32, h int32) (bm windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateCompatibleBitmap.Addr(), 3, uintptr(dc), uintptr(w), uintptr(h))
	bm = windows.Handle(r0)
	if bm == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func CreateCompatibleDC(dc windows.Handle) (mdc windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateCompatibleDC.Addr(), 1, uintptr(dc), 0, 0)
	mdc = windows.Handle(r0)
	if mdc == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DeleteDC(dc windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteDC.Addr(), 1, uintptr(dc), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func DeleteObject(obj windows.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteObject.Addr(), 1, uintptr(obj), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SelectObject(dc windows.Handle, obj windows.Handle) (oldobj windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procSelectObject.Addr(), 2, uintptr(dc), uintptr(obj), 0)
	oldobj = windows.Handle(r0)
	if oldobj == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetPixel(dc windows.Handle, x int32, y int32, color uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetPixel.Addr(), 4, uintptr(dc), uintptr(x), uintptr(y), uintptr(color), 0, 0)
	if r1 == ^uintptr(0) {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func Shell_NotifyIcon(message uint32, data *NotifyIconData) (err error) {
	r1, _, e1 := syscall.Syscall(procShell_NotifyIconW.Addr(), 2, uintptr(message), uintptr(unsafe.Pointer(data)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
