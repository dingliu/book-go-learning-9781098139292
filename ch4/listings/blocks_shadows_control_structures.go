package main

import (
	"fmt"
	"math/rand"
)

func shadowing() {
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10
}

func shadowConstant() {
	fmt.Println(true) // true
	true := 10
	fmt.Println(true) // 10
}

func ifElse() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

func ifElseVariableScope() {
	// n scoped to the condition, if, and else
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

func completeFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forWithoutInit() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forWithoutIncrement() {
	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}

func conditionOnlyFor() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

// func infiniteFor() {
// 	for {
// 		fmt.Println("Hello")
// 	}
// }

func whileLoop() {
	i := 1
	for {
		fmt.Println("Hello")
		i++
		if !(i < 10) {
			break
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

func forRangeLoop() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}
}

func forRangeCopyValue() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v = v * 2
	}
	fmt.Println(evenVals)
}

func labelForLoops() {
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

func switchStatement() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			fmt.Println(word, "is exactly the right length:", size)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}
}

func switchBreakForWithLabel() {
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is borring")
		}
	}
}

func blankSwitch() {
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word.")
		case wordLen > 10:
			fmt.Println(word, "is a long word.")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}

func fizzBuzzWithSwitch() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	shadowing()
	shadowConstant()
	ifElse()
	ifElseVariableScope()
	completeFor()
	forWithoutInit()
	forWithoutIncrement()
	conditionOnlyFor()
	// infiniteFor()
	whileLoop()
	fizzBuzz()
	forRangeLoop()
	forRangeCopyValue()
	labelForLoops()
	switchStatement()
	switchBreakForWithLabel()
	blankSwitch()
	fizzBuzzWithSwitch()
}
