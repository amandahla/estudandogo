package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

// teste removendo o '*' de Person
func (p *Person) Describe() string {
	return fmt.Sprintf("Name: %s Age: %d", p.name, p.age)
}

func (p *Person) ChangeName(newName string)  {
	p.name = newName
}

func main() {
	p1 := Person{"Bob", 40}
	fmt.Println(p1.Describe())
	p1.ChangeName("Peter")
    fmt.Println(p1.Describe())
}