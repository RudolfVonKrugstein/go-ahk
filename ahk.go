// +build windows

package go_ahk

import (
	"syscall"
	"time"
)

type AutoHotKeyDLL struct {
	dll *syscall.LazyDLL
	currentScript *AHKScript
}

func NewAutoHotKeyDLL() (*AutoHotKeyDLL, error) {
	return &AutoHotKeyDLL{
		dll: syscall.NewLazyDLL("./AutoHotkey.dll"),
	}, nil
}

func isError(err error) bool {
	return err != nil && err.Error() != "The operation completed successfully."
}

func (dll* AutoHotKeyDLL) RunScript(script *AHKScript) error {
	// Are we already running?
	if dll.currentScript != nil {
		dll.WaitForScript()
	}
	_, _, err := dll.dll.NewProc("ahktextdll").Call(script.GetUIntPtr(), uintptr(0), uintptr(0))
	if isError(err) {
		return err
	}
	dll.currentScript = script
	return nil
}

func (dll* AutoHotKeyDLL) IsScriptRunning() (bool, error) {
	if dll.currentScript == nil {
		return false, nil
	}
	res, _, err := dll.dll.NewProc("ahkReady").Call()
	if isError(err) {
		return false, err
	}
	return res != 0, nil
}

func (dll* AutoHotKeyDLL) WaitForScript() error {
	running := true
	var err error
	for ;running; {
		running, err = dll.IsScriptRunning()
		if isError(err) {
			return err
		}
		if running {
			time.Sleep(time.Millisecond * 100)
		}
	}
	return nil
}
