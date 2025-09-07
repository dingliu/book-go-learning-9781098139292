package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
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

func makePtr[T any](t T) *T {
	return &t
}

func useMakePtr() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	p := person{
		FirstName:  "Pat",
		MiddleName: makePtr("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println(p) // {Pat 0xc000012110 Peterson}
}

func failToUpdateNilPtrParam() {
	failedUpdate := func(g *int) {
		x := 10
		g = &x
	}

	var f *int // f is a nil pointer
	failedUpdate(f)
	fmt.Println(f) // prints nil
}

func mustDerefPtrParam() {
	failedUpdate := func(ptrX *int) {
		x2 := 20
		ptrX = &x2
	}
	update := func(ptrX *int) {
		*ptrX = 20
	}

	x := 10
	failedUpdate(&x)
	fmt.Println(x) // 10
	update(&x)
	fmt.Println(x) // 20
}

func readFileWithBuffer() error {
	processData := func([]byte) {
		// do something
	}
	filename := os.Args[0]
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	data := make([]byte, 100)
	for {
		count, err := f.Read(data)
		processData(data[:count])
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
	}
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

	slog.Info("Use generic helper")
	useMakePtr()

	slog.Info("failed to update nil pointer")
	failToUpdateNilPtrParam()

	slog.Info("Must deferencing pointer to update value")
	mustDerefPtrParam()
}
