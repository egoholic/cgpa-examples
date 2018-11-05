package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("INterface Flags:", i.Flags.String())
		fmt.Println("Interfaces MTU", i.MTU)
		fmt.Println("Inteface hardware address:", i.HardwareAddr)
		fmt.Println()
	}
}
