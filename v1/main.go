package main

import "fmt"

func main() {

	// methods - special type of func
	// func associated with a invocation and a var

	// interfaces

	// oop concept is diff from other langs

	fmt.Println(isEven(4))

	var zmi mi = 5
	fmt.Println(zmi.isOdd())

	// polymorphism concept
	// interface var
	var s Speaker

	// assign dog to speaker
	s = Dog{}
	s.Speak()

	// assign human to speaker
	s = Human{}
	s.Speak()

	// one version of the function that works for many types - generic programming
	printItem("hello")
	printItem(3.14)
	// in generics when using map the key should be comparable type not any
	// + operator (etc.) cannot be used with any or struct or comparable

	// type constraint interface restricts what types can be used with a generic type parameter
	fmt.Println(add(2.5, 3.5))

}

type addable interface {
	int | float64
}

func add[T addable](a, b T) T {
	return a + b
}

// any will accept all types
func printItem[T any](item T) {
	fmt.Println(item)
}

// Any type that has a Speak() method is a Speaker
type Speaker interface {
	Speak()
}

type Dog struct{}

type Human struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func (h Human) Speak() {
	fmt.Println("Hello!")
}

// func
// bool is return type
// ans := isEven(i)
func isEven(i int) bool {
	return i%2==0
}

// method
// (i mi) is reciever
// executes in context of var i
// tight coupling between func and type
// ans := mi.isEven()
type mi int
func (i mi) isOdd() bool {
	return int(i)%2!=0
}