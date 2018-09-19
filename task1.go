package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f := os.Stdin
	s := bufio.NewScanner(f)
	var n float64
	var err error
	counter := 0
	sum := 0.0
	var text string

	for s.Scan() {
		text = s.Text()
		if text == "STOP" {
			break
		}

		n, err = strconv.ParseFloat(text, 64)
		if err == nil {
			counter++
			sum = sum + n
			fmt.Println("Please enter a number:")
		} else {
			fmt.Println(text, "is not a number!")
		}
	}

	fmt.Println("SUM:", sum)
	fmt.Println("AVG:", sum/float64(counter))
}
