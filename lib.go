package main

/*
#include <wchar.h>
#include <stdint.h>
#include <stdbool.h>
*/
import "C"

const DUMMYHTTP_URL = "http://localhost:8080"

// definition of class names, as UTF-16 strings
// from "A"
var classNames = []C.wchar_t{0x0041, 0x0000}

//export GetClassNames
func GetClassNames() *C.wchar_t {
	logToConsoleGo("GetClassNames was called")
	return &(classNames[0])
}

//export GetClassObject
func GetClassObject(className *C.wchar_t, component **C.void) C.long {
	classNameStr := WCharToString(className)
	switch classNameStr {
	case "A":
		return CreateComponentA(component)
	default:
		return 0
	}
}

//export DestroyObject
func DestroyObject(component **C.void) C.long {
	logToConsoleGo("DestroyObject was called")
	return 1
}

//export SetPlatformCapabilities
func SetPlatformCapabilities(capabilities *C.long) C.long {
	logToConsoleGo("SetPlatformCapabilities was called")
	return 1
}

func main() {}
