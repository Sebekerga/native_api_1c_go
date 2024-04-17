package main

/*
#include <wchar.h>
*/
import "C"
import (
	"unicode/utf16"
	"unsafe"
)

func WCharToString(wc *C.wchar_t) string {
	if wc == nil {
		return ""
	}
	const max = 0xffff
	p := (*[max]uint16)(unsafe.Pointer(wc))
	s := p[:max]
	for i, v := range s {
		if v == 0 {
			s = s[0:i]
			break
		}
	}
	return string(utf16.Decode(s))
}

func WCharFromString(s string, trailingZero bool) []C.wchar_t {
	runes := utf16.Encode([]rune(s))
	if trailingZero {
		runes = append(runes, 0)
	}
	cRunes := make([]C.wchar_t, len(runes))
	for i, r := range runes {
		cRunes[i] = C.wchar_t(r)
	}
	return cRunes
}

func WCharPtrFromString(s string, trailingZero bool) *C.wchar_t {
	runes := WCharFromString(s, trailingZero)
	return (*C.wchar_t)(unsafe.Pointer(&runes[0]))
}
