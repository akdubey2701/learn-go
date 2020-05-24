/*package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (a Animal) Name() string {
	return a.name
}
func (a Animal) Eat() string {
	return a.food
}
func (a Animal) Move() string {
	return a.locomotion
}
func (a Animal) Speak() string {
	return a.sound
}
func (a *Animal) Init(n, f, l, s string) {
	a.sound = s
	a.name = n
	a.locomotion = l
	a.food = f
}
func main() {
	var A_list [3]Animal
	var name string
	var ch string
	A_list[0].Init("cow", "grass", "walk", "moo")
	A_list[1].Init("bird", "worms", "fly", "peep")
	A_list[2].Init("snake", "mice", "slither", "hsss")
	fmt.Print("Name:")
	fmt.Scan(&name)
	fmt.Print("Choice(“eat”, “move”, or “speak”):")
	fmt.Scan(&ch)
	ch = strings.ToLower(ch)
	name = strings.ToLower(name)
	for i := 0; i < 3; i++ {
		if strings.Compare(name, A_list[i].Name()) == 0 {
			if strings.Compare(ch, "eat") == 0 {
				fmt.Println(A_list[i].Eat())
			}
			if strings.Compare(ch, "move") == 0 {
				fmt.Println(A_list[i].Move())
			}
			if strings.Compare(ch, "speak") == 0 {
				fmt.Println(A_list[i].Speak())
			}

		}
	}
	fmt.Print(A_list)
}
*/
//  golang course
//  Month 2	Week 3  Assignment 1
/*
	You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, sound string
}

func (an Animal) Eat() {
	fmt.Println(an.food)
}

func (an Animal) Move() {
	fmt.Println(an.locomotion)
}

func (an Animal) Speak() {
	fmt.Println(an.sound)
}

func main() {

	m := map[string]Animal{
		"cow": Animal{
			"grass",
			"walk",
			"moo",
		},
		"bird": Animal{
			"worms",
			"fly",
			"peep",
		},
		"snake": Animal{
			"mice",
			"slither",
			"hsss",
		}}

	for {
		fmt.Print(">  ")
		br := bufio.NewReader(os.Stdin)
		data, _, _ := br.ReadLine()

		task := strings.TrimSpace(string(data))
		strArray := strings.Split(task, " ")
		//	fmt.Printf("%T  %T   %T\n", strArray, data, task)

		animalName := strArray[0]
		animalTask := strArray[1]

		if animalTask == "eat" {
			m[animalName].Eat()
		} else if animalTask == "move" {
			m[animalName].Move()
		} else if animalTask == "speak" {
			m[animalName].Speak()
		}
	}
}
*/
package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (eat *Animal) Eat() {
	fmt.Println(eat.food)
}

func (move *Animal) Move() {
	fmt.Println(move.locomotion)
}

func (speak *Animal) Speak() {
	fmt.Println(speak.noise)
}

func main() {

	var animal, action string
	var err error

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {

		fmt.Printf("> ")

		_, err = fmt.Scan(&animal)
		if err != nil {
			fmt.Println(err)
		}

		_, err = fmt.Scan(&action)
		if err != nil {
			fmt.Println(err)
		}

		if animal == "cow" {
			switch action {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		}

		if animal == "bird" {
			switch action {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		}

		if animal == "snake" {
			switch action {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}
		}
	}
}
