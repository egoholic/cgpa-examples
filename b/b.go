package b

// use: $ go install cgpa-examples/b
// to install the package

import (
	"cgpa-examples/a"
	"fmt"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("FromB()")
	a.FromA()
}
