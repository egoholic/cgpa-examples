package main

import (
	"a" //should be installed
	"b" // should be installed
	"fmt"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	a.FromA()
	b.FromB()
}
