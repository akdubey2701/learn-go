/*package main

import (
	"fmt"
)

func Swap(x, y int) (int, int) {
	return y, x
}
func BubbleSort(sli [10]int, n int) [10]int {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = Swap(sli[j], sli[j+1])
			}
		}
	}
	return sli
}
func main() {
	var a [10]int
	fmt.Printf("Enter Size Of Array:")
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter Next Element:")
		fmt.Scan(&a[i])
	}
	a = BubbleSort(a, n)
	fmt.Println(a[:n])
}
*/
package main

import (
	"fmt"
)

func main() {

	numbers := make([]int, 10)
	fmt.Println("\nEnter sequence of 10 space seperated integers: \n")
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("Entered numbers -> %+v \n\n", numbers)
	fmt.Printf("Sorting numbers using bubble sort........\n\n")
	BubbleSort(numbers)
	fmt.Printf("Sorted numbers -> %+v \n\n", numbers)
}

// Takes in slice of int and sorts ascending
func BubbleSort(numbers []int) {
	size := len(numbers)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			first := numbers[j]
			next := numbers[j+1]
			// compare curr and next number if curr>next swap
			if Compare(first, next) {
				Swap(numbers, j)
			}
		}
	}
}

// Compares if first integer greate than second one
func Compare(num1, num2 int) bool {
	return num1 > num2
}

// Swaps position of two ints at index pos i and i+1 in a slice
func Swap(ints []int, i int) {
	ints[i], ints[i+1] = ints[i+1], ints[i]
}
