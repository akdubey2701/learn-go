/*package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	name string
	addr string
}

func main() {
	var name string
	var addr string
	fmt.Println("Enter Name:")
	fmt.Scan(&name)
	fmt.Println("Enter Address:")
	fmt.Scan(&addr)

	var usr User
	usr.name = name
	usr.addr = addr
	barr, err := json.Marshal(usr)
	if err != nil {
		fmt.Println(barr)

	}
	fmt.Println("JSON:", barr)
}*/
//  golang course
//  Month 1	Module 4    Week 4  Assignment 1
/*
	Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string) //		create map
	fmt.Println("Enter your name")
	var name string
	fmt.Scan(&name)
	m["name"] = name
	fmt.Println("Enter your Address")
	var address string
	fmt.Scan(&address)
	m["address"] = address

	//	create json using Marshal
	data, err := json.Marshal(m)

	if err != nil {
		fmt.Println("Error Occured")
	} else {
		fmt.Println(string(data)) //	converting json into string
	}
}
