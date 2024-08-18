package main

import "fmt"

func main() {
	ex1() // 20 20
	ex2() // 42 42
	ex3() // 0 -2147483648 0
}

func ex1() {
	var x int = 20
	var f float64 = float64(x)
	fmt.Println(x, f)
}

func ex2() {
	const value = 42
	var i int = value
	var f float64 = value
	fmt.Println(i, f)
}

func ex3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
