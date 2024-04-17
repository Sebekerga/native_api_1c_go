package main

/*
#include <stdbool.h>
*/
import "C"
import "unsafe"

//export _register_extension_as
func _register_extension_as(self *C.struct_AddIn, extension_name *C.wchar_t) C.bool {
	extension_name = WCharPtrFromString("A", true)
	return true
}

//export _get_n_props
func _get_n_props(self *C.struct_AddIn) C.long {
	logToConsoleGo("Get number of properties")
	return 1
}

//export _find_prop
func _find_prop(self *C.struct_AddIn, prop_name *C.wchar_t) C.long {
	logToConsoleGo("Find property")
	return 0
}

//export _get_prop_name
func _get_prop_name(self *C.struct_AddIn, prop_index C.long, lang_index C.long) *C.wchar_t {
	logToConsoleGo("Get property name")
	// TODO: Implement _get_prop_name
	return nil
}

//export _get_prop_val
func _get_prop_val(self *C.struct_AddIn, prop_index C.long, value unsafe.Pointer) C.bool {
	logToConsoleGo("Get property value")
	// TODO: Implement _get_prop_val
	return false
}

//export _set_prop_val
func _set_prop_val(self *C.struct_AddIn, prop_index C.long, value unsafe.Pointer) C.bool {
	logToConsoleGo("Set property value")
	// TODO: Implement _set_prop_val
	return false
}

//export _is_prop_readable
func _is_prop_readable(self *C.struct_AddIn, prop_index C.long) C.bool {
	logToConsoleGo("Is property readable")
	// TODO: Implement _is_prop_readable
	return false
}

//export _is_prop_writable
func _is_prop_writable(self *C.struct_AddIn, prop_index C.long) C.bool {
	// TODO: Implement _is_prop_writable
	return false
}

//export _get_n_methods
func _get_n_methods(self *C.struct_AddIn) C.long {
	// TODO: Implement _get_n_methods
	return 0
}

//export _find_method
func _find_method(self *C.struct_AddIn, method_name *C.wchar_t) C.long {
	// TODO: Implement _find_method
	return 0
}

//export _get_method_name
func _get_method_name(self *C.struct_AddIn, method_index C.long, lang_index C.long) *C.wchar_t {
	// TODO: Implement _get_method_name
	return nil
}

//export _get_n_params
func _get_n_params(self *C.struct_AddIn, method_index C.long) C.long {
	// TODO: Implement _get_n_params
	return 0
}

//export _get_param_default_val
func _get_param_default_val(self *C.struct_AddIn, method_index C.long, param_index C.long, value unsafe.Pointer) C.bool {
	// TODO: Implement _get_param_default_val
	return false
}

//export _call_as_proc
func _call_as_proc(self *C.struct_AddIn, method_index C.long, params unsafe.Pointer) C.bool {
	// TODO: Implement _call_as_proc
	return false
}

//export _call_as_func
func _call_as_func(self *C.struct_AddIn, method_index C.long, params unsafe.Pointer, result unsafe.Pointer) C.bool {
	// TODO: Implement _call_as_func
	return false
}
