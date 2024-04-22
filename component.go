package main

/*
#include <wchar.h>
#include <stdint.h>
#include <stdbool.h>

typedef struct ComponentA
{
    int counter;
} ComponentA;

#define ADDIN_TYPE ComponentA
#include <ffi.h>


static bool AllocMemory(AddInInterface *self, void **memory, uint32_t size)
{
    return self->memory_manager->vtable->alloc_memory(self->memory_manager, memory, size);
}

static void FreeMemory(AddInInterface *self, void **memory)
{
    self->memory_manager->vtable->free_memory(self->memory_manager, memory);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

type AttachType = C.enum_AttachType
type PlatformCapabilities = C.enum_PlatformCapabilities
type IAddInDefBase = C.struct_IAddInDefBase

var ErrMemoryAllocationFailed = errors.New("memory allocation failed")

func (addIn *AddInInterface) AllocateMemory(size int) (unsafe.Pointer, error) {
	pointerToData := unsafe.Pointer(nil)
	ok := C.AllocMemory(addIn, &pointerToData, C.uint32_t(size))
	if !ok {
		return nil, ErrMemoryAllocationFailed
	}
	Logger.Printf("Allocated memory of size %d\n", size)

	return pointerToData, nil
}

// ComponentA is a struct that represents a component
type ComponentA = C.struct_ComponentA

func CreateComponentA(pointerToAddInPointer **C.void) C.long {
	componentWrapper := C.CreateGenericComponent()
	C.InitComponent(componentWrapper, &ComponentA{counter: 0})
	*pointerToAddInPointer = (*C.void)(unsafe.Pointer(componentWrapper))
	return 1
}

//export _register_extension_as
func _register_extension_as(self **LanguageExtenderVTable, extension_name *C.wchar_t) C.bool {
	Logger.Printf("Registering extension as '%s'\n", WCharToString(extension_name))
	return true
}

//export _get_n_props
func _get_n_props(self **LanguageExtenderVTable) C.long {
	Logger.Println("Getting number of properties")
	return 1
}

//export _find_prop
func _find_prop(self **LanguageExtenderVTable, prop_name *C.wchar_t) C.long {
	Logger.Printf("Finding property '%s'\n", WCharToString(prop_name))
	return 0
}

//export _get_prop_name
func _get_prop_name(self **LanguageExtenderVTable, prop_index C.long, lang_index C.long) *C.wchar_t {
	Logger.Printf("Getting property name for index %d\n", prop_index)
	return nil
}

//export _get_prop_val
func _get_prop_val(self **LanguageExtenderVTable, prop_index C.long, value *PlatformVar) C.bool {
	Logger.Printf("Getting property value for index %d\n", prop_index)
	err := value.SetString(GetParent(self), "Hello, world!")
	if err != nil {
		Logger.Printf("Failed to set string value: %v\n", err)
		return false
	}
	return true
}

//export _set_prop_val
func _set_prop_val(self **LanguageExtenderVTable, prop_index C.long, value *PlatformVar) C.bool {
	Logger.Printf("Setting property value for index %d\n", prop_index)
	return true
}

//export _is_prop_readable
func _is_prop_readable(self **LanguageExtenderVTable, prop_index C.long) C.bool {
	return true
}

//export _is_prop_writable
func _is_prop_writable(self **LanguageExtenderVTable, prop_index C.long) C.bool {
	return true
}

//export _get_n_methods
func _get_n_methods(self **LanguageExtenderVTable) C.long {
	return 0
}

//export _find_method
func _find_method(self **LanguageExtenderVTable, method_name *C.wchar_t) C.long {
	return -1
}

//export _get_method_name
func _get_method_name(self **LanguageExtenderVTable, method_index C.long, lang_index C.long) *C.wchar_t {
	return nil
}

//export _get_n_params
func _get_n_params(self **LanguageExtenderVTable, method_index C.long) C.long {
	return 0
}

//export _get_param_default_val
func _get_param_default_val(self **LanguageExtenderVTable, method_index C.long, param_index C.long, value unsafe.Pointer) C.bool {
	return true
}

//export _call_as_proc
func _call_as_proc(self **LanguageExtenderVTable, method_index C.long, params unsafe.Pointer) C.bool {
	return true
}

//export _call_as_func
func _call_as_func(self **LanguageExtenderVTable, method_index C.long, params unsafe.Pointer, result unsafe.Pointer) C.bool {
	return true
}
