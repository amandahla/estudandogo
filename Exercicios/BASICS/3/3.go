/*
3) Altere o programa criado anteriormente para que o número máximo seja obtido
por meio de uma função getRandomNumber(n) em que n é o valor máximo a ser gerado.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber(n int) int {
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(n)
}

func main() {
	sum := 0

	maximumNumber := getRandomNumber(100)

	for i := 1; i < maximumNumber; i++ {
		isMultipleOf3 := i%3 == 0
		isMultipleOf5 := i%5 == 0

		if isMultipleOf3 || isMultipleOf5 {
			sum += i
		}
	}

	fmt.Printf("Maximum number is %v\n", maximumNumber)

	fmt.Println(sum)
}
