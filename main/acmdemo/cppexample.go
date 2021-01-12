package acmdemo

import "fmt"

// #cgo LDFLAGS: -L../../cpp/speedUplib/lib/windows/static/ -lwindowsStatic
// #include "../../cpp/speedUplib/lib/bridge.h"
import "C"

func MainCpp() {
	fmt.Println("==============================================================")
	fmt.Println("C++")
	fmt.Println("==============================================================")
	fmt.Println(C.Increment(3))
}
