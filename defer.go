package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ") // remembers current i
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ") // uses latest i (i = 0)
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ") // remembers current i
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
