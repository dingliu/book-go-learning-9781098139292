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

func pointerTypes() {
	x := 10
	var ptrX *int = &x // *int is the pointer type of ptrX
	fmt.Println(*ptrX) // 10
}

func newPtrVar() {
	x := new(int)         // x is a pointer of int
	fmt.Println(x == nil) // false
	fmt.Println(*x)       // 0
}

func ptrStruct() {
	type foo struct {
		bar string
	}

	ptrFoo := &foo{
		bar: "test",
	}
	fmt.Println(*ptrFoo) // {test}
}

func ptrPrimitiveLiteral() {
	var x string
	ptrX := &x
	fmt.Println(*ptrX)
}

func main() {
	slog.Info("first pointer function")
	ptrFuncOne()

	slog.Info("Address operator and dereferencing operator")
	addressAndDereferencingOperator()

	slog.Info("Pointer types")
	pointerTypes()

	slog.Info("Built-in new function")
	newPtrVar()

	slog.Info("Pointer of a struct literal")
	ptrStruct()

	slog.Info("Pointer of a string, primitive type")
	ptrPrimitiveLiteral()
}
