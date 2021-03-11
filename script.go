package go_ahk

import "unsafe"

/**
 An AHK Script.
 It must exists
 */
type AHKScript struct {
	script []byte
}

func NewAHKScript(s string) *AHKScript {
	// Convert to wide char string
	bytes := make([]byte, (len(s) + 1) * 2)
	for i := 0; i < len(s); i++ {
	bytes[i * 2] = s[i]
	bytes[i * 2 + 1] = 0
	}
	bytes[len(s) * 2] = 0
	bytes[len(s) * 2 + 1] = 0

	return &AHKScript{
		bytes,
	}
}

func (s *AHKScript) GetUIntPtr() uintptr {
	return uintptr(unsafe.Pointer(&s.script[0]))
}
