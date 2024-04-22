package main

/*
#include <wchar.h>
#include <stdint.h>
#include <stdbool.h>

*/
import "C"
import (
	"unsafe"
)

type MemoryManager = C.struct_MemoryManager
type Connection = C.struct_Connection
type AddInInterface = C.struct_AddInInterface
type LanguageExtenderVTable = C.struct_LanguageExtenderVTable

func GetParent(lang_ext **LanguageExtenderVTable) *AddInInterface {
	Logger.Printf("Getting AddIn pointer from LanguageExtenderVTable: %v", lang_ext)

	pointerSize := unsafe.Sizeof(lang_ext)
	offset := 1
	Logger.Printf("Applying offset: %d, pointer size: %d", offset, pointerSize)

	addInPointer := (*AddInInterface)(unsafe.Pointer(uintptr(unsafe.Pointer(lang_ext)) - uintptr(offset)*pointerSize))

	Logger.Printf("AddIn pointer: %v", unsafe.Pointer(addInPointer))

	return addInPointer
}

// definition of class names, as UTF-16 strings
// from "A"
var classNamesGo = "A"
var classNames = [...]C.wchar_t{0x0041, 0x0}

//export GetClassNames
func GetClassNames() *C.uint16_t {
	Logger.Println("- - - - - - Platform is loading AddIn - - - - - -")
	Logger.Printf("GetClassNames was called, returning: %v", classNamesGo)

	wcharSize := unsafe.Sizeof(classNames[0])

	arrayPointer := unsafe.Pointer(&classNames)
	for i := range classNames {
		pointer := unsafe.Pointer(uintptr(arrayPointer) + uintptr(i)*wcharSize)

		Logger.Printf(
			"#%d: %v = %#x, %v = %#x",
			i,
			(*C.uint8_t)(unsafe.Pointer(uintptr(pointer))),
			*(*C.uint8_t)(unsafe.Pointer(uintptr(pointer))),
			(*C.uint8_t)(unsafe.Pointer(uintptr(pointer)+1)),
			*(*C.uint8_t)(unsafe.Pointer(uintptr(pointer) + 1)),
		)
	}

	Logger.Printf("Array pointer: %v", arrayPointer)
	return (*C.uint16_t)(arrayPointer)
}

//export GetClassObject
func GetClassObject(name *C.uint16_t, component **C.void) C.long {
	Logger.Println("- - - - Platform is loading AddIn object - - - -")

	// return 0
	Logger.Printf("GetClassObject was called, className: %s", WCharToString(name))

	classNameStr := WCharToString(name)
	switch classNameStr {
	case "A":
		Logger.Printf("Matched class %s, creating component A", classNameStr)
		return CreateComponentA((**C.void)(unsafe.Pointer(component)))
	default:
		Logger.Printf("No match for class %s", classNameStr)
		return 1
	}
}

//export DestroyObject
func DestroyObject(component **C.void) C.long {
	Logger.Println("DestroyObject was called")
	return 1
}

//export SetPlatformCapabilities
func SetPlatformCapabilities(capabilities *PlatformCapabilities) PlatformCapabilities {
	Logger.Println("SetPlatformCapabilities was called")
	return 1
}

//export GetAttachType
func GetAttachType() AttachType {
	Logger.Println("- - - - Platform is getting attach type - - - - -")
	return 1
}

func main() {}
