package main

import "fmt"

const const_x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const const_z = 20 * 20

func constExample() {
	fmt.Println(const_x)
	fmt.Println(const_z)

	// const_x = const_x + 1 // this won't compile
	// CONST_Z = "bye"       // this wont compile as well
}
