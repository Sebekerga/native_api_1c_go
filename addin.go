package main

/*
#include <addin.h>
*/
import "C"

func CreateComponent(pointerToAddInPointer **C.struct_AddIn, component C.struct_Component) C.long {
	*pointerToAddInPointer = C.CreateAddIn()
	return 1
}
