package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

type myFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func simulateNamedParams(opts myFuncOpts) error {
	fmt.Println("first name is:", opts.FirstName)
	return errors.New("something")
}

func addto(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, val := range vals {
		out = append(out, base+val)
	}
	return out
}

func divAndRemainder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func f1(str string) int {
	return len(str)
}

func f2(str string) int {
	var total int = 0
	for _, ch := range str {
		total += int(ch)
	}
	return total
}

func demoAnonymousFunc() {
	f := func(j int) {
		fmt.Println("in anonymous func", j)
	}
	for i := 0; i < 5; i++ {
		f(i)
	}

	func(j int) {
		fmt.Println("in anonymous func", j)
	}(99)
}

func main() {
	slog.Info("demonstrating functions")
	fmt.Println(div(4, 2))

	slog.Info("simulating named params")
	simulateNamedParams(myFuncOpts{
		FirstName: "Joe",
		LastName:  "Doe",
		Age:       80,
	})

	slog.Info("demonstrating variadic functions")
	fmt.Println(addto(3))
	fmt.Println(addto(3, 2))
	fmt.Println(addto(3, 2, 4, 6, 8))

	a := []int{4, 3}
	fmt.Println(addto(3, a...))
	fmt.Println(addto(3, []int{4, 3}...))

	slog.Info("demonstrating multiple return values")
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	slog.Info("demonstrating function as values")
	var f func(string) int
	f = f1
	fmt.Println(f("hello"))
	f = f2
	fmt.Println(f("hello"))

	type myOpFuncType func(int, int) int
	type myOpFuncMapType map[string]myOpFuncType
	var myOpFunc myOpFuncType
	var myOpFuncMap myOpFuncMapType

	myOpFunc = func(a int, b int) int {
		return a + b
	}
	fmt.Println(myOpFunc(3, 4))
	myOpFuncMap = map[string]myOpFuncType{
		"add": func(a int, b int) int {
			return a + b
		},
		"sub": func(a int, b int) int {
			return a - b
		},
	}
	fmt.Println(myOpFuncMap["add"](3, 4))

	slog.Info("demonstrating anonymous functions")
	demoAnonymousFunc()
}
