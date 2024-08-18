package main

import "fmt"

func main() {
	compareArrays()
	zeroSliceIsNil()
}

func compareArrays() {
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y)
}

func zeroSliceIsNil() {
	var x []int
	fmt.Println(x == nil)
}
