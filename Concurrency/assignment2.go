//go run -race assignment2.go
/*A race condition is an undesirable situation that occurs when a device or system attempts to perform two or more operations at the same time, but because of the nature of the device or system, the operations must be done in the proper sequence to be done correctly.*/
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// function f is called directly , function f and another anynomous function  are interleaved to show race condition
func main() {
	f("call executed directly")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("go routine interleaved")
	time.Sleep(time.Second)
	fmt.Println("done")
}
