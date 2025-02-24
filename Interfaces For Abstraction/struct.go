/*package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
	// type_of_animal string
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
func (a *Animal) Init(name, type_of_animal string) {
	a.name = name
	if strings.Compare(type_of_animal, "cow") == 0 {
		a.food = "grass"
		a.locomotion = "walk"
		a.noise = "moo"

	} else if strings.Compare(type_of_animal, "bird") == 0 {
		a.food = "worms"
		a.locomotion = "fly"
		a.noise = "peep"
	} else if strings.Compare(type_of_animal, "snake") == 0 {
		a.food = "mice"
		a.locomotion = "slither"
		a.noise = "hsss"
	}
	// a.type_of_animal=type_of_animal
}
func main() {

	var comm, animal, action string
	list := make([]Animal, 0, 20)
	var temp Animal
	var err error

	/*	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
		bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
		snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {

		fmt.Printf("> ")

		_, err = fmt.Scan(&comm)
		if err != nil {
			fmt.Println(err)
		}
		_, err = fmt.Scan(&animal)
		if err != nil {
			fmt.Println(err)
		}

		_, err = fmt.Scan(&action)
		if err != nil {
			fmt.Println(err)
		}

		if strings.Compare(comm, "newanimal") == 0 {
			temp.Init(animal, action)
			list = append(list, temp)
			fmt.Println("created it!")
		} else if strings.Compare(comm, "query") == 0 {
			for i := 0; i < len(list); i++ {
				temp = list[i]
				if strings.Compare(temp.name, animal) == 0 {
					switch action {
					case "eat":
						temp.Eat()
					case "move":
						temp.Move()
					case "speak":
						temp.Speak()
					}
				}
			}
		}
	}
}
*/
package main

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (ani animal) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani animal) Move() {
	fmt.Println(ani.locomotion)
	return
}

func (ani animal) Speak() {
	fmt.Println(ani.noise)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}
	var genralAni animalInterface
	for {
		var command, requestAni, requestType string
		fmt.Print(">")
		fmt.Scan(&command, &requestAni, &requestType)
		if command == "query" {
			genralAni = animalMap[requestAni]
			switch requestType {
			case "eat":
				genralAni.Eat()
			case "move":
				genralAni.Move()
			case "speak":
				genralAni.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[requestAni] = animalMap[requestType]
			fmt.Println("Created it!")
		}
	}
}
