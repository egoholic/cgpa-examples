package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	b := []byte{}
	buffer := bytes.NewBuffer(b)
	buffer.Write([]byte("This is"))
	fmt.Fprintf(buffer, " a string!\n")
	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stdout)

	buffer.Reset()
	buffer.Write([]byte("Mastering Go!"))
	buffer.Write([]byte(buffer.String()))
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 3)
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}
}
