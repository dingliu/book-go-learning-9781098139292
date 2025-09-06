package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sort"
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

func closureDemo() {
	i := 0
	f := func() {
		fmt.Println(i)
		i++
	}
	f()
	fmt.Println("in closureDemo, i is", i)
}

func sortingSlice() {
	type person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []person{
		{"Pat", "Paterson", 30},
		{"Joe", "Doe", 80},
		{"Jane", "Doe", 70},
	}
	fmt.Println("before sorting:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("after sorting by last name:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("after sorting by Age:", people)
}

func makeMultiplier(factor int) func(int) int {
	return func(val int) int {
		return factor * val
	}
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
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

	slog.Info("demonstrating closures")
	closureDemo()

	slog.Info("demonstrating sorting with anonymous functions")
	sortingSlice()

	slog.Info("demonstrating function factories")
	times2 := makeMultiplier(2)
	times3 := makeMultiplier(3)
	fmt.Println("3 times 2 is", times2(3))
	fmt.Println("4 times 3 is", times3(4))

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)
	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}
