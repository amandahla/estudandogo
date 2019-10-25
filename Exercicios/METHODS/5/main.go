package main

import (
	"fmt"
)

type Animal interface {
	Walk()
}

type Dog struct {
	Name string
}

func (d *Dog) Walk() {
	if d == nil {
		fmt.Println("value <nil> received")
		return
	}
	fmt.Println("walking")
}

func (d Dog) Bark() {
	fmt.Println("barking")
}

func DoSomething(a Animal) {
	a.Walk()
}

func main() {
	thor := &Dog{"Thor"}
	thor = nil
	DoSomething(thor)
}
