package main

import (
	"fmt"
	"slices"
)

func main() {
	compareArrays()
	zeroSliceIsNil()
	compareSlices()
	appendSlices()
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
	// s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // true
	fmt.Println(slices.Equal(y, z)) // false
	// fmt.Println(slices.Equal(x, s)) // won't compile
}

func appendSlices() {
	var x []int
	x = append(x, 10)
	x = append(x, 5, 6, 7)
	y := []int{10, 20, 30}
	x = append(x, y...)
	fmt.Println(x) // [10, 5, 6, 7, 10, 20, 30]
}
