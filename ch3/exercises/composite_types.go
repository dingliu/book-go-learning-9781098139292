package main

import "fmt"

func grettingSlices() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[4:]
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(greetings)
}

func main() {
	grettingSlices()
}
