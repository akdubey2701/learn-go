package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func inputList() []int {
	list := make([]int, 0, 4)

	for {
		fmt.Print("Enter a number, or 'X' or 'x' to exit: ")
		var s string
		fmt.Scanf("%s", &s)
		s = strings.ToLower(s)

		if s == "x" {
			break
		}

		v, err := strconv.Atoi(s)

		if err != nil {
			fmt.Printf("%s, try again\n", err)
			continue
		}

		list = append(list, int(v))

	}

	return list
}

// This is from a previous assigment, so I thought I'd used it instead of sort.Ints
func bubbleSort(list []int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("sorting sublist %v\n", list)

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				swap(list, j)
			}
		}
	}
}

func swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

func removeMin(sublists [][]int) int {
	mini := -1
	minv := int(math.MaxInt32)
	for i, sublist := range sublists {
		if len(sublist) > 0 && sublist[0] < minv {
			mini = i
			minv = sublist[0]
		}
	}

	// remove the smallest value from its respective sublist
	if mini != -1 {
		sublists[mini] = sublists[mini][1:]
	}

	return minv
}

func main() {

	list := inputList()
	sublists := make([][]int, 4)

	// chop up list into as evenly sized pieces as possible
	sublists[0] = list[:len(list)/4]
	sublists[1] = list[len(list)/4 : len(list)/2]
	sublists[2] = list[len(list)/2 : 3*len(list)/4]
	sublists[3] = list[3*len(list)/4:]

	var wg sync.WaitGroup
	wg.Add(4)
	go bubbleSort(sublists[0], &wg)
	go bubbleSort(sublists[1], &wg)
	go bubbleSort(sublists[2], &wg)
	go bubbleSort(sublists[3], &wg)
	wg.Wait()

	sortedList := make([]int, 0, len(list))
	for range list {
		v := removeMin(sublists)
		sortedList = append(sortedList, v)
	}

	fmt.Printf("sortedList: %v\n", sortedList)
}
