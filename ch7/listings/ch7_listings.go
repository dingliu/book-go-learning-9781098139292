package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

type Score int
type Converter func(string) Score
type TeamScores map[string]Score

type HighScore Score

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func (e Employee) String() string {
	return fmt.Sprintf("%s %s, age %d", e.FirstName, e.LastName, e.Age)
}

type Manager struct {
	Employee
	Reports []Employee
}

// func demoNoInheritance() {
// 	var i int = 100        // assign untyped literal value, ok
// 	var s Score = 100      // assign untyped literal value, ok
// 	var hs HighScore = 200 // assign untyped literal value, ok
// 	hs = s                 // compilation error
// 	s = i                  // compilation error
// 	s = Score(i)           // assign with type conversion, ok
// 	hs = HighScore(s)      // assign with type conversion, ok
// }

func embeddingDemo() {
	bob := Manager{
		Employee: Employee{
			FirstName: "Bob",
			LastName:  "Bobson",
			Age:       35,
		},
		Reports: []Employee{},
	}
	fmt.Println(bob.Age)    // promoted field
	fmt.Print(bob.String()) // promoted method
}

func methodDemo() {
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	output := p.String()
	fmt.Println(output)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("Total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Stringer interface {
	String() string
}

type Incrementer interface {
	Increment()
}

// func demoInterface() {
// 	var myStringer Stringer
// 	var myIncrementer Incrementer

// 	pointerCounter := &Counter{}
// 	valueCounter := Counter{}

// 	myStringer = pointerCounter
// 	myStringer = valueCounter

// 	myIncrementer = pointerCounter // ok
// 	myIncrementer = valueCounter   // cannot use valueCounter (variable
// 	// of struct type Counter) as Incrementer value in assignment:
// 	// Counter does not implement Incrementer (method Increment has
// 	// pointer receiver)compilerInvalidIfaceAssign
// }

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("In doUpdateWrong", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("In doUpdateRight", c.String())
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Social
	Spam
	Advertisements
)

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

type Activity int

const (
	_ Activity = iota // 0 indicates uninitialized
	Sleeping
	Walking
	Running
)

// (implicit) implementation of the interface required by the client
type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// process data
	fmt.Println(data)
	return data
}

// (explicit) declaration of the required interface on Client side
type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program(data string) {
	c.L.Process(data)
}

func putAltogether() {
	data := "data"
	c := Client{
		L: LogicProvider{},
	}
	c.Program(data)
}

// decorator pattern
func process(r io.Reader) error {
	fmt.Println("")
	return errors.New("some")
}

func decoratorDemo1(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()
	return process(r)
}

func decoratorDemo2(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()
	gz, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	return process(gz)
}

func interfaceEmbedding() {
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Closer interface {
		Close() error
	}

	type ReadCloser interface {
		Reader
		Closer
	}
}

func interfaceAndNil() {
	var pointerCounter *Counter
	fmt.Println(pointerCounter == nil)
	var incrementer Incrementer
	fmt.Println(incrementer == nil)
	incrementer = pointerCounter
}

// interface comparison
type Doubler interface {
	Double()
}

type DoubleInt int

func (di *DoubleInt) Double() {
	*di = *di * 2
}

type DoubleIntSlice []int

func (dis DoubleIntSlice) Double() {
	for i := range dis {
		dis[i] *= 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

// type assertion
type MyInt int

func typeAssertionDemo() error {
	var i any
	var mine MyInt = 20
	i = mine

	i2, ok := i.(int)
	if !ok {
		return fmt.Errorf("unexpected type")
	}
	fmt.Println(i2 + 1)
	return nil
}

func typeSwitchDemo(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("j is nill")
	case int:
		fmt.Printf("j is an int: %v\n", j)
	case string:
		fmt.Printf("j is a string: %s\n", j)
	case io.Reader:
		fmt.Println("j is of type io.Reader")
	case bool, rune:
		fmt.Println("i is either a bool or rune, so j is of type any")
	default:
		fmt.Println("no idea what i is, so j is of type any")
	}
}

func main() {
	slog.Info("Method demo")
	methodDemo()

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	slog.Info("value receiver and pointer receiver")
	var c2 Counter
	doUpdateWrong(c2)
	fmt.Println(c2.String())
	doUpdateRight(&c2)
	fmt.Println(c2.String())

	slog.Info("methods on nil instances")
	var it *IntTree
	it.Insert(5)
	it.Insert(3)
	it.Insert(10)
	it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))

	slog.Info("Methods as functions")
	myAdder := Adder{start: 10}

	var f func(int) int = myAdder.AddTo
	fmt.Println(f(10))
	slog.Info("Iota")
	fmt.Println(Field1, Field2, Field3, Field4, Field5)

	slog.Info("Embedding demo")
	embeddingDemo()

	slog.Info("Implicit interface demo")
	putAltogether()

	slog.Info("Compare interfaces")
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis DoubleIntSlice = DoubleIntSlice{1, 2, 3}
	// var dis2 DoubleIntSlice = DoubleIntSlice{1, 2, 3}

	DoublerCompare(&di, &di2)
	DoublerCompare(&di, dis)
	// DoublerCompare(dis, dis2) triggers panic
	// panic: runtime error: comparing uncomparable type main.DoubleIntSlice

	slog.Info("Type assertions")
	typeAssertionDemo()
}
