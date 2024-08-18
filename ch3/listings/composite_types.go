package main

import "fmt"

func main() {
	compareArrays()
}

func compareArrays() {
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y)
}
