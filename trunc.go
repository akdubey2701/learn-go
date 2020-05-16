package main

/*package main

import "fmt"

func main() {
	var x float32
	fmt.Println("Enter no:")
	fmt.Scan(&x)
	y:=int(x)
	fmt.Println(y)
}
//truncation program

package main
import "fmt"

func main() {
	var x float64
	var y int
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&x)
	y=int(x)
	fmt.Print("Truncated verion is: ")
	fmt.Print(y)
}*/
//package main

import "fmt"

func main() {
	var userFloat float64

	fmt.Printf("Enter float to truncate: ")

	// Get input
	_, err := fmt.Scan(&userFloat)
	if err != nil {
		// Error due to invalid input
		fmt.Println("Invalid float")
	} else {
		// Valid input so truncate
		fmt.Println(int(userFloat))
	}
}
