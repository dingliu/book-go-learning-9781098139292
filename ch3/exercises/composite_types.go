package main

import "fmt"

func grettingSlices() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[4:]
	fmt.Println(first)     // [Hello Hola]
	fmt.Println(second)    // [Hola नमस्कार こんにちは]
	fmt.Println(third)     // [Привіт]
	fmt.Println(greetings) // [Hello Hola नमस्कार こんにちは Привіт]
}

func fourthRuneInSlice() {
	message := "Hi 👩 and 👨"
	var rs []rune = []rune(message)
	fmt.Println(string(rs[3])) // 👩
}

func employeeStruct() {
	type employee struct {
		firstName string
		lastName  string
		id        int
	}

	a := employee{
		"John",
		"Doe",
		1,
	}
	b := employee{
		firstName: "Jane",
		lastName:  "Doe",
		id:        2,
	}
	var c employee
	c.firstName = "Bob"
	c.lastName = "Smith"
	c.id = 3

	fmt.Println(a) // {John Doe 1}
	fmt.Println(b) // {Jane Doe 2}
	fmt.Println(c) // {Bob Smith 3}
}

func main() {
	grettingSlices()
	fourthRuneInSlice()
	employeeStruct()
}
