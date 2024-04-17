package main

/*
#include <wchar.h>
#include <stdint.h>
#include <addin.h>

*/
import "C"
import (
	"fmt"
)

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
func GetClassObject(className *C.wchar_t, component **C.struct_AddIn) C.long {
	classNameStr := WCharToString(className)
	logToConsoleGo(fmt.Sprintf("GetClassObject was called, className: %s", classNameStr))

	logToConsoleGo("Creating Component in Go")
	creatingResult := CreateComponent(component, C.struct_Component{})
	logToConsoleGo("Component created in Go")

	return creatingResult
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
