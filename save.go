package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	s := []byte("Data to write\n")

	// 1st method
	f1, err := os.Create("tmp/f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	// 2nd method
	f2, err := os.Create("tmp/f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	// 3rd method
	f3, err := os.Create("tmp/f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	// 4th method
	f4 := "tmp/f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 5th method
	f5, err := os.Create("tmp/f5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("wrote %d bytes\n", n)
}
