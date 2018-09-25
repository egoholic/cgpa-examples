package main

import (
	"fmt"
)

var Arr = [3]int{1, 2, 3}

func changeArray(a [3]int) [3]int {
	a[0] = 100
	return a
}

func main() {
	fmt.Println("Original array:", Arr)
	fmt.Println("ChangeArray() result:", changeArray(Arr))
	fmt.Println("Original array:", Arr)
}
