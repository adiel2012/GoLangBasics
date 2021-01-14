package acmdemo

//#cgo LDFLAGS:  -L../../cpp/speedUplib/lib/windows/dynamic/ -lwindowsDynamic
// #include "../../cpp/speedUplib/lib/facade/bridge.h"
import "C"
import (
	"fmt"
	"unsafe"
)

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
}
