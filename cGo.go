package main

//#include <stdio.h>
//void callC() {
//  printf("Calling C code!\n");
//}
import "C"

// Looks like C package is a kinda trick.
// There should be no empty string above the `import "C"` line,
// cause you will get error:
//     ./cGo.go:13:2: could not determine kind of name for C.callC

import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}
