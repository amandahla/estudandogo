/*
5) Altere o programa criado anteriormente para:
- Ao invés de exibir somente a soma, deverá exibir "Sum is"
- Tanto para a mensagem da soma quanto para a mensagem retornada pela função getRandomNumber,
o programa deverá checar a variável printInPortuguese.
Se o valor dela for true, exibirá "Soma é" e "Número máximo é".
Se for false, exibirá "Sum is" e "Maximum number is"
Dica: onde a variável printInPortuguese deve ser declarada para ser utilizada tanto
pela function main quanto pela getRandomNumber?
Saída de exemplo:
O número máximo é 34
Soma é 258
*/
package main

import (
	"fmt"
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
}
