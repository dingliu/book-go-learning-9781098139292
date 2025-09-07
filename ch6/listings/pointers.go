package main

import (
	"fmt"
	"log/slog"
)

func ptrFuncOne() {
	var x int32 = 10
	var y bool = true
	ptrX := &x
	ptrY := &y
	var ptrZ *string

	fmt.Println(ptrX)
	fmt.Println(ptrY)
	fmt.Println(ptrZ)
}

func addressAndDereferencingOperator() {
	x := "Hello"
	ptrX := &x
	fmt.Println(ptrX)
	fmt.Println(*ptrX)
	y := *ptrX + " World!"
	fmt.Println(y)
}

func main() {
	slog.Info("first pointer function")
	ptrFuncOne()
	slog.Info("Address operator and dereferencing operator")
	addressAndDereferencingOperator()
}
