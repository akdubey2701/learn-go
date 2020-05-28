package main

import (
	"fmt"
)

func p(a, b, c int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
func main() {
	go p(1, 2, 3)
	go p(4, 5, 6)

}

//go run -race assignment2.go
/*A race condition is an undesirable situation that occurs when a device or system attempts to perform two or more operations at the same time, but because of the nature of the device or system, the operations must be done in the proper sequence to be done correctly.*/
