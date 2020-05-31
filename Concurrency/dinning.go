package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	name        int
	left, right *Chopstick
}

func (p Philosopher) eat(mayIEat chan int, doneEating chan int, wg *sync.WaitGroup) {
	fmt.Println("\t\t\t--- philosopher", p.name, "created")

	for i := 0; i < maxEats; i++ {
		<-mayIEat
		p.left.Lock()
		p.right.Lock()

		fmt.Println("starting to eat", p.name)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println("finishing eating", p.name)

		p.right.Unlock()
		p.left.Unlock()
		doneEating <- p.name
	}

	fmt.Println("\t\t\t--- philosopher", p.name, "done")

	wg.Done()
}

func createPhilosophers() []*Philosopher {
	sticks := make([]*Chopstick, totPhilosophers)
	for i := 0; i < totPhilosophers; i++ {
		sticks[i] = new(Chopstick)
	}

	philos := make([]*Philosopher, totPhilosophers)
	for i := 0; i < totPhilosophers; i++ {
		philos[i] = &Philosopher{i + 1, sticks[i], sticks[(i+1)%totPhilosophers]}
	}
	return philos
}

func host(eatSlots []chan int, doneSlots []chan int, group *sync.WaitGroup) {
	eat := 0
	totEats := totPhilosophers * maxEats

	eatSlots[0] <- startEating
	eatSlots[1] <- startEating
	for {
		select {
		case p1 := <-doneSlots[0]:
			fmt.Println("\t\t\t--- eat slot 1 available, was", p1)
			eatSlots[0] <- startEating
			eat++
		case p2 := <-doneSlots[1]:
			fmt.Println("\t\t\t--- eat slot 2 available, was", p2)
			eatSlots[1] <- startEating
			eat++
		}

		if eat == totEats {
			break
		}
	}

	fmt.Println("\t\t\t--- host done")
	group.Done()
}

const totPhilosophers = 5
const maxEats = 3
const startEating = 1

func main() {

	philos := createPhilosophers()

	wg := new(sync.WaitGroup)

	eatSlots := make([]chan int, 2)
	eatSlots[0] = make(chan int, totPhilosophers/2+1)
	eatSlots[1] = make(chan int, totPhilosophers/2+1)

	doneSlots := make([]chan int, 2)
	doneSlots[0] = make(chan int, totPhilosophers/2+1)
	doneSlots[1] = make(chan int, totPhilosophers/2+1)

	for i := 0; i < totPhilosophers; i++ {
		s := i % 2
		go philos[i].eat(eatSlots[s], doneSlots[s], wg)
		wg.Add(1)
	}

	go host(eatSlots, doneSlots, wg)
	wg.Add(1)

	wg.Wait()
}
