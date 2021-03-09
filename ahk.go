package go_ahk

import (
	"syscall"
	"time"
	"unsafe"
)

type AutoHotKeyDLL struct {
	dll *syscall.LazyDLL
}

func NewAutoHotKeyDLL() (*AutoHotKeyDLL, error) {
	return AutoHotKeyDLL{
		dll: syscall.NewLazyDLL("./AutoHotkey.dll"),
	}, nil
}

func stringToAhkStringBytes(s string) []byte {
	bytes := make([]byte, (len(s) + 1) * 2)
	for i := 0; i < len(s); i++ {
		bytes[i * 2] = s[i]
		bytes[i * 2 + 1] = 0
	}
	bytes[len(s) * 2] = 0
	bytes[len(s) * 2 + 1] = 0
	return bytes
}

func isError(err error) bool {
	return err != nil && err.Error() != "The operation completed successfully."
}

func (dll* AutoHotKeyDLL) RunScript(script string) error {
	bytes := stringToAhkStringBytes(script)

	_, _, err := dll.dll.NewProc("ahktextdll").Call(uintptr(unsafe.Pointer(&bytes[0])), uintptr(0), uintptr(0))
	if isError(err) {
		return err
	}
	return nil
}

func (dll* AutoHotKeyDLL) IsScriptRunning() (bool, error) {
	res, _, err := dll.dll.NewProc("ahkReady").Call()
	if isError(err) {
		return false, err
	}
	return res != 0, nil
}

func (dll* AutoHotKeyDLL) WaitForScript() error {
	running := true
	for ;running; {
		running, err := dll.IsScriptRunning()
		if isError(err) {
			return err
		}
		if running {
			time.Sleep(time.Millisecond * 100)
		}
	}
	return nil
}
