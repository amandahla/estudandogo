/*
2) Altere o programa criado anteriormente para que o número máximo seja aleatório.
Dica: "math/rand" e não esqueça de usar "rand.Seed(time.Now().UTC().UnixNano())"
para não retornar sempre o mesmo valor
Imprima na tela o valor do número máximo com a frase "Maximum number is " e depois a soma.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0

	rand.Seed(time.Now().UTC().UnixNano())

	maximumNumber := rand.Intn(1000)

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
