package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type person struct {
	FName string
	LName string
}

func main() {
	maxLength := 20
	var p person
	fmt.Printf("Enter File name:")
	var x string
	fmt.Scan(&x)
	content, err := ioutil.ReadFile(x)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	words := strings.Fields(text)
	for i := 1; i < len(words); i += 2 {
		p.FName = words[i-1]
		if len(p.FName) > maxLength {
			p.FName = p.FName[:maxLength]
		}
		p.LName = words[i]
		if len(p.LName) > maxLength {
			p.LName = p.LName[:maxLength]
		}
		fmt.Println((i+1)/2, ":First Name:", p.FName, "\t\tLast Name:", p.LName)
	}
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Print("enter filename: ")
	fmt.Scanln(&filename)
	f, _ := os.Open(filename)

	sli := []*name{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())

		n1 := new(name)
		n := strings.Split(scanner.Text(), " ")
		n1.fname, n1.lname = n[0], n[1]

		sli = append(sli, n1)

	}

	for i := range sli {
		n := sli[i]
		fmt.Println(n.fname + " " + n.lname)
	}

}
*/
