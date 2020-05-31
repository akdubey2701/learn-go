package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("\nProgram runs 2 goroutines forever showing how output differs unpredictabley in runs!\n")
	fmt.Println("X starts as '1' intention is to add 1 to make it '2' and print but routines unpredictable\n")
	for {
		x := 1
		wg := sync.WaitGroup{}
		wg.Add(2)
		time.Sleep(2 * time.Second)
		go addOne(&x, &wg)
		go printX(&x, &wg)
		wg.Wait()
	}
}

func addOne(x *int, wg *sync.WaitGroup) {
	*x += 1
	wg.Done()
}

func printX(x *int, wg *sync.WaitGroup) {
	fmt.Println(*x)
	wg.Done()
}
