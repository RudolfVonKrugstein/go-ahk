// +build linux

// For non windows platforms, no operation library

package go_ahk

type AutoHotKeyDLL struct {
}

func NewAutoHotKeyDLL() (*AutoHotKeyDLL, error) {
	return &AutoHotKeyDLL{}, nil
}


func (dll* AutoHotKeyDLL) RunScript(script AHKScript) error {
	return nil
}

func (dll* AutoHotKeyDLL) IsScriptRunning() (bool, error) {
	return false, nil
}

func (dll* AutoHotKeyDLL) WaitForScript() error {
	return nil
}
