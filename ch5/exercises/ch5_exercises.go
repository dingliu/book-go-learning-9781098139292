package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
)

// exercise 1

func add(i, j int) (int, error) { return i + j, nil }
func sub(i, j int) (int, error) { return i - j, nil }
func mul(i, j int) (int, error) { return i * j, nil }
func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("Division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// exercise 2
func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	size := 0
	data := make([]byte, 100)
	for {
		n, err := f.Read(data)
		size += n
		if err != nil {
			if err == io.EOF {
				slog.Info("end of file")
				break
			}
			return size, err
		}
	}
	return size, nil
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("Invalid expression:", exp)
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operator:", op)
		}
		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println("Got error:", err)
		}
		fmt.Println(result)
	}
	fileSize, err := fileLen("file")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileSize)
}
