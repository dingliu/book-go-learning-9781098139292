package main

import "fmt"

var (
	x    int
	y        = 20
	z    int = 30
	d, e     = 40, "hello"
	f, g string
)

func main() {
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)

	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println(sum3, sum4)

	untypedLiterals()
	varDeclareWithinFunc()
}

func untypedLiterals() {
	var x float64 = 10
	var y float64 = 200.3 * 5
	fmt.Println(x, y)
}

func varDeclareWithinFunc() {
	x := 10
	x, y := 20, "hello"
	fmt.Println(x, y)
}
