/*package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("Enter Sring:")
	fmt.Scan(&x)
	a := "a"
	t1 := strings.Contains(x, a)
	//fmt.Println(i)
	i := "i"
	t2 := strings.HasPrefix(x, i)
	n := "n"
	t3 := strings.HasSuffix(x, n)
	if t1 && t2 && t3 {
		fmt.Printf("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userString string

	fmt.Printf("Enter string: ")

	// Get user input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userString = scanner.Text()
	}

	// Trim trailing spaces
	userString = strings.TrimSpace(userString)

	// Convert to lower to handle mixed case
	userString = strings.ToLower(userString)

	// Check for presence of `i`, `a` and `n`
	if strings.HasPrefix(userString, "i") && strings.Contains(userString, "a") && strings.HasSuffix(userString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
