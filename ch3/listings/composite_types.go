package main

import (
	"fmt"
	"maps"
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

func slicingSlices() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
}

func slicesOverlapMemory() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x: ", x) // x:  [x y z d]
	fmt.Println("y: ", y) // y:  [x y]
	fmt.Println("z: ", z) // z:  [y z d]
}

func confusingSliceAppend() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y)) // 4 4
	y = append(y, "z")
	fmt.Println("x: ", x) // x:  [a b z d]
	fmt.Println("y: ", y) // y:  [a b z]
}

func evenMoreConfusingSliceAppend() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x: ", x) // x: [a, b, i, j, y]
	fmt.Println("y: ", y) // y: [a, b, i, j, y]
	fmt.Println("z: ", z) // z: [i, j, y]
}

func fullSliceExpressions() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z)) // 5 2 2
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x: ", x) // x: [a, b, c, d, x]
	fmt.Println("y: ", y) // y: [a, b, i, j, k] (created a new slice in memory)
	fmt.Println("z: ", z) // z: [c, d, y] (created a new slice in memory)
}

func convertSliceToPointerToArray() {
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)

	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println(xSlice)        // [10, 20, 3, 4]
	fmt.Println(xArrayPointer) // &[10, 20, 3, 4]
}

func extraAndSliceStrings() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b) // 116
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2) // o t
	fmt.Println(s3) // Hello
	fmt.Println(s4) // there
}

func sliceStringWithEmoji() {
	var s string = "Hello 🌞"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2)     // o �
	fmt.Println(s3)     // Hello
	fmt.Println(s4)     // 🌞
	fmt.Println(len(s)) // 10
}

func convertIntToString() {
	var x int = 65
	var s string = string(x)
	fmt.Println(s)
}

// ./composite_types.go:163:17: conversion from int to string yields a
// string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)

func convertStringToSlices() {
	var s string = "Hello 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs) // [72 101 108 108 111 32 240 159 140 158]
	fmt.Println(rs) // [72 101 108 108 111 32 127774]
}

func declareMaps() {
	var nilMap map[string]int
	totalWins := map[string]int{}
	teams := map[string][]string{
		"Orcas":   []string{"John", "Jane"},
		"Lions":   []string{"Sarah", "Peter"},
		"Kittens": []string{"Waldo", "Raul"},
	}
	fmt.Println(nilMap)
	fmt.Println(totalWins)
	fmt.Println(teams)
}

func makeMap() {
	ages := make(map[int][]string, 10)
	fmt.Println(len(ages)) // 0
}

func commaOkIdiom() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true

	v, ok = m["world"]
	fmt.Println(v, ok) // 0 true

	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0 false
}

func deleteAndClearMap() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	delete(m, "hello")
	fmt.Println(m) // map[world:0]
	clear(m)
	fmt.Println(m) // map[]
}

func compareMapEquality() {
	m1 := map[string]int{
		"hello": 5,
		"world": 0,
	}

	m2 := map[string]int{
		"hello": 5,
		"world": 0,
	}

	fmt.Println(maps.Equal(m1, m2)) // true
}

func asSet() {
	intSet := map[int]bool{}
	values := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range values {
		intSet[v] = true
	}
	fmt.Println(len(values), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

func main() {
	compareArrays()
	zeroSliceIsNil()
	compareSlices()
	appendSlices()
	sliceCapacity()
	makeSlices()
	clearSlice()
	slicingSlices()
	slicesOverlapMemory()
	confusingSliceAppend()
	evenMoreConfusingSliceAppend()
	fullSliceExpressions()
	convertSliceToPointerToArray()
	extraAndSliceStrings()
	sliceStringWithEmoji()
	convertIntToString()
	convertStringToSlices()
	declareMaps()
	makeMap()
	commaOkIdiom()
	deleteAndClearMap()
	compareMapEquality()
	asSet()
}
