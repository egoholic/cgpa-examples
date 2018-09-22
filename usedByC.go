package main

import "C"

//To compile use:
//   $ go build -o usedByC.o -buildmode=c-shared usedByC.go

import (
	"fmt"
)

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {

}
