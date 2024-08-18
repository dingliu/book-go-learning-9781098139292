package main

import (
	"fmt"
	"slices"
)

func main() {
	compareArrays()
	zeroSliceIsNil()
	compareSlices()
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

func compareSlices() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // true
	fmt.Println(slices.Equal(y, z)) // false
	// fmt.Println(slices.Equal(x, s)) // won't compile
}
