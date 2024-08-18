package main

import (
	"fmt"
	"slices"
)

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

func sliceCapacity() {
	var x []int
	fmt.Println(x, len(x), cap(x)) // [] 0 0
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x)) // [10] 1 1
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x)) // [10 20] 2 2
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x)) // [10 20 30] 3 4
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40] 4 4
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x)) // [10 20 30 40 50] 5 8
}

func makeSlices() {
	x := make([]int, 5)
	y := make([]int, 5, 10)
	z := make([]int, 0, 10)
	fmt.Println(x, y, z)
}

func clearSlice() {
	s := []string{"first", "second", "third"}
	clear(s)
	fmt.Println(s, len(s))
}

func main() {
	compareArrays()
	zeroSliceIsNil()
	compareSlices()
	appendSlices()
	sliceCapacity()
	makeSlices()
	clearSlice()
}
