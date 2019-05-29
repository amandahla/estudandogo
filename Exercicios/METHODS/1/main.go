package main

import (
	"fmt"
	"time"
)

const layoutISO = "2006-01-02"

type Person struct {
	name string
	birthday string
}

func (p Person) Age() string {
	start, _ := time.Parse(layoutISO, p.birthday)
	end := time.Now()
	difference := end.Sub(start)

	return fmt.Sprintf("%d", int64(difference.Hours()/24)/365)
}

func main() {
	p1 := Person{"Amanda", "1985-12-25"}
	fmt.Println(p1.Age())
}