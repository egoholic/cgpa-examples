package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

// C is a virtual Go package that just tells go build to preprocess its
// input file using the cgo tool before the Go compiler processes the file.

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.cHello()
	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Volodymyr!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)
	fmt.Println("All perfectly done!")
}
