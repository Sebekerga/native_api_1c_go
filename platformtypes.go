package main

/*
#include <platformtypes.h>
*/
import "C"
import (
	"time"
	"unsafe"
)

type PlatformValue = C.union_PlatformValue
type PlatformVar = C.struct_PlatformVar

type PlatformWCharString = C.struct_WCharString

func (v *PlatformVar) SetBool(value bool) {
	Logger.Println("Setting bool value")
	*(*C.bool)(unsafe.Pointer(&v.value)) = C.bool(value)
	v.ty = C.PlatTypeBool
}

func (v *PlatformVar) SetInt(value int32) {
	Logger.Println("Setting int value")
	*(*C.int)(unsafe.Pointer(&v.value)) = C.int(value)
	v.ty = C.PlatTypeInt32
}

func (v *PlatformVar) SetFloat(value float64) {
	Logger.Println("Setting float value")
	*(*C.double)(unsafe.Pointer(&v.value)) = C.double(value)
	v.ty = C.PlatTypeFloat64
}

func (v *PlatformVar) SetTime(value time.Time) {
	Logger.Printf("Setting time value (GO) %v\n", value)
	platformTime := C.struct_PlatformTime{
		sec:   C.int(value.Second()),
		min:   C.int(value.Minute()),
		hour:  C.int(value.Hour()),
		mday:  C.int(value.Day()),
		mon:   C.int(value.Month()),
		year:  C.int(value.Year() - 1900),
		wday:  C.int(value.Weekday()),
		yday:  C.int(value.YearDay()),
		isdst: 0,
	}
	Logger.Printf("Setting time value (1C) %v\n", platformTime)

	*(*C.struct_PlatformTime)(unsafe.Pointer(&v.value)) = platformTime
	v.ty = C.PlatTypeTime
}

func (v *PlatformVar) SetString(addIn *AddInInterface, value string) error {
	Logger.Printf("Setting string value %v", value)
	logPointerToConsole((*C.void)(unsafe.Pointer(addIn)))
	wcharString := WCharFromString(value, false)
	targetLen := len(wcharString) * int(unsafe.Sizeof(wcharString[0]))
	Logger.Printf("WChar length: %d\n", len(wcharString))
	Logger.Printf("Target length: %d\n", targetLen)

	allocatedPointer, err := addIn.AllocateMemory(targetLen)
	if err != nil {
		Logger.Printf("Failed to allocate memory: %v\n", err)
		return err
	}

	for i, wchar := range wcharString {
		Logger.Printf("%s - WChar #%d: %v\n", string(value[i]), i, wchar)
		pointerToWChar := unsafe.Pointer(uintptr((allocatedPointer)) + uintptr(i)*unsafe.Sizeof(wchar))
		*(*C.wchar_t)(pointerToWChar) = wchar
	}

	*(*PlatformWCharString)(unsafe.Pointer(&v.value)) = PlatformWCharString{
		str: (*C.wchar_t)(unsafe.Pointer(allocatedPointer)),
		len: C.uint(len(wcharString)),
	}
	v.ty = C.PlatTypeWStr

	return nil
}
