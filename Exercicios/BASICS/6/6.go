/*
6) Altere o programa criado anteriomente para exibir a raiz quadrada de sum
com a mensagem "Sqrt is". Dica: a função Sqrt espera um parâmetro do tipo float64.
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var printInPortuguese = true

func getRandomNumber(n int) (randomNumber int, message string) {
	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber = rand.Intn(n)
	message = fmt.Sprintf("Maximum number is %v", randomNumber)
	if printInPortuguese {
		message = fmt.Sprintf("Número máximo é %v", randomNumber)
	}
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

	messageSum := fmt.Sprintf("Sum is %v", sum)
	if printInPortuguese {
		messageSum = fmt.Sprintf("Soma é %v", sum)
	}
	fmt.Println(messageSum)

	fmt.Printf("Sqrt is %v\n", math.Sqrt(float64(sum)))
}
