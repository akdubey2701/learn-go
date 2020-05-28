package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	size := 20
	slice := generateSlice(size)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", Sort_func(slice, size), "\n")
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
func Sort_func(slice []int, size int) []int {

	l0 := 0
	l1 := int(size / 4)
	l2 := int(size/4 + 1)
	l3 := int(size / 2)
	l4 := int(size/2 + 1)
	l5 := int(3 * size / 4)
	l6 := int(3*size/4 + 1)
	l7 := size - 1
	go sort(slice, l0, l1)
	go sort(slice, l2, l3)
	go sort(slice, l4, l5)
	go sort(slice, l6, l7)

	sort(slice, 0, size-1)

	return slice
}
func sort(sli []int, s, e int) []int {
	for i := s; i <= e; i++ {
		for j := s; j <= e-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	return sli
}
