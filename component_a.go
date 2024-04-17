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
*/
import "C"
import "unsafe"

// ComponentA is a struct that represents a component
type ComponentA = C.struct_ComponentA
type AddInInterface = C.struct_AddInInterface

func CreateComponentA(pointerToAddInPointer **C.void) C.long {
	componentWrapper := C.CreateGenericComponent()
	C.InitComponent(componentWrapper, &ComponentA{counter: 0})
	*pointerToAddInPointer = (*C.void)(unsafe.Pointer(componentWrapper))
	return 1
}

//export _register_extension_as
func _register_extension_as(self *AddInInterface, extension_name *C.wchar_t) C.bool {
	return true
}

//export _get_n_props
func _get_n_props(self *AddInInterface) C.long {
	return 0
}

//export _find_prop
func _find_prop(self *AddInInterface, prop_name *C.wchar_t) C.long {
	return -1
}

//export _get_prop_name
func _get_prop_name(self *AddInInterface, prop_index C.long, lang_index C.long) *C.wchar_t {
	return nil
}

//export _get_prop_val
func _get_prop_val(self *AddInInterface, prop_index C.long, value unsafe.Pointer) C.bool {
	return true
}

//export _set_prop_val
func _set_prop_val(self *AddInInterface, prop_index C.long, value unsafe.Pointer) C.bool {
	return true
}

//export _is_prop_readable
func _is_prop_readable(self *AddInInterface, prop_index C.long) C.bool {
	return true
}

//export _is_prop_writable
func _is_prop_writable(self *AddInInterface, prop_index C.long) C.bool {
	return true
}

//export _get_n_methods
func _get_n_methods(self *AddInInterface) C.long {
	return 0
}

//export _find_method
func _find_method(self *AddInInterface, method_name *C.wchar_t) C.long {
	return -1
}

//export _get_method_name
func _get_method_name(self *AddInInterface, method_index C.long, lang_index C.long) *C.wchar_t {
	return nil
}

//export _get_n_params
func _get_n_params(self *AddInInterface, method_index C.long) C.long {
	return 0
}

//export _get_param_default_val
func _get_param_default_val(self *AddInInterface, method_index C.long, param_index C.long, value unsafe.Pointer) C.bool {
	return true
}

//export _call_as_proc
func _call_as_proc(self *AddInInterface, method_index C.long, params unsafe.Pointer) C.bool {
	return true
}

//export _call_as_func
func _call_as_func(self *AddInInterface, method_index C.long, params unsafe.Pointer, result unsafe.Pointer) C.bool {
	return true
}
