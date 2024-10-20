package main

import (
	"fmt"
	"math/rand"
)

func exerciseOne() {
	s := []int{}

	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(100))
	}

	for _, v := range s {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind.")
		}
	}

}

func exerciseTwo() {
	total := 0
	for i := 0; i < 10; i++ {
		// total := total + i
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}

func main() {
	exerciseOne()
	exerciseTwo()
}
