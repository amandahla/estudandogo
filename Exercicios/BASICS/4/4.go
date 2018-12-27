/*
4) Altere o programa criado anteriormente para que a função getRandomNumber(n) retorne o número e a mensagem a ser exibida "Maximum number is".
Dica: message = fmt.Sprintf("Maximum number is %v\n", randomNumber)
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber(n int) (randomNumber int, message string) {
	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber = rand.Intn(n)
	message = fmt.Sprintf("Maximum number is %v", randomNumber)
	return
}

func main() {
	sum := 0

	maximumNumber, message := getRandomNumber(100)

	for i := 1; i < maximumNumber; i++ {
		isMultipleOf3 := i%3 == 0
		isMultipleOf5 := i%5 == 0

		if isMultipleOf3 || isMultipleOf5 {
			sum += i
		}
	}

	fmt.Println(message)

	fmt.Println(sum)
}
