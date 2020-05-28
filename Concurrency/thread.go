package main

import (
	"fmt"
	"time"
)

func p() {
	fmt.Println("5")
}
func main() {
	go p()

	fmt.Println("1")
	fmt.Println("2")
	go p()
	fmt.Println("3")
	go p()
	time.Sleep(10 * time.Millisecond)

}
