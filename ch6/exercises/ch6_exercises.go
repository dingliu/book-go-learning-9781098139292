package main

import (
	"fmt"
	"time"
)

// Exercise 1
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makePerson(fn string, ln string, age int) Person {
	return Person{
		FirstName: fn,
		LastName:  ln,
		Age:       age,
	}
}

func makePersonPtr(fn string, ln string, age int) *Person {
	return &Person{
		FirstName: fn,
		LastName:  ln,
		Age:       age,
	}
}

// Exercise 2
func updateSlice(sli []string, str string) {
	lasPos := len(sli) - 1
	sli[lasPos] = str
	fmt.Println("The updated slice is", sli)
}

func growSlice(sli []string, str string) {
	sli = append(sli, str)
	fmt.Println("The updated slice is", sli)
}

func main() {
	// Exercise 1
	fn := "John"
	ln := "Doe"
	age := 33

	p := makePerson(fn, ln, age)
	ptrP := makePersonPtr(fn, ln, age)

	fmt.Println(p)
	fmt.Println(ptrP)

	// Exercise 2
	s := []string{"hello", "world", "foo"}
	fmt.Println("Before update slice", s)
	updateSlice(s, "bar")
	fmt.Println("After update slice", s)

	fmt.Println("Before grow slice", s)
	growSlice(s, "bar")
	fmt.Println("After grow slice", s)

	before := time.Now()
	for i := 0; i < 10_000_000; i++ {
		_ = Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       99,
		}
	}
	after := time.Now()
	duration := after.Sub(before)
	fmt.Println("The time took for create 10,000,000 Person struct is", duration)
}
