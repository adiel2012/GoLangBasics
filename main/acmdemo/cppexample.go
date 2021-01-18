package acmdemo

//#cgo LDFLAGS:  -L../../cpp/speedUplib/lib/windows/dynamic/ -lwindowsDynamic
// #include "../../cpp/speedUplib/lib/facade/bridge.h"
/*
#include <stdlib.h>
#include <stdio.h>

int callOnMeGo_cgo(int in); // Forward declaration.

*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export displayNumber
func displayNumber(in int) int {
	//fmt.Printf("Go.callOnMeGo(): called with arg = %d\n", in)
	if in > 10 {
		return 0
	}
	fmt.Println(in)
	return 1
}

type A struct {
	ptr unsafe.Pointer
}

func NewA(value int) A {
	var a A
	a.ptr = C.New_A(C.int(value))
	return a
}

func (a A) Free() {
	C.Release_A(a.ptr)
}

func (a A) value() int {
	return int(C.GetValue_A(a.ptr))
}

func MainCpp() {
	fmt.Println("==============================================================")
	fmt.Println("C++")
	fmt.Println("==============================================================")
	fmt.Println(C.Increment(3))
	a := NewA(37)
	defer a.Free() // The Go analog to C++'s RAII
	fmt.Println(a.value())

	fmt.Printf("Go.main(): calling C function with callback to us\n")
	C.GenerateFibonacci((C.callback_fcn)(unsafe.Pointer(C.callOnMeGo_cgo)))

}
