/*
1) Resolver o problema 1 do Projeto Euler.
Se listamos todos os números naturais inferiores a 10 que são múltimos de 3 ou 5,
temos 3,5,6,e 9. A soma desses valores é 23.
O programa deve imprimir na tela a soma dos múltiplos de 3 ou 5 inferiores a 1000.
*/
package main

import (
	"fmt"
	//"unsafe"
	//"time"
)

func main() {
	//var sum int // variables declared without an explicit initial value are given their zero value (numeric = 0)
	var sum = 0
	//sum := 0 // short declaration

	//fmt.Println("Size in bytes of int:", unsafe.Sizeof(sum))

	//start := time.Now()

	for i := 1; i < 1000; i++ {
		isMultipleOf3 := i%3 == 0
		isMultipleOf5 := i%5 == 0

		//if i%3 == 0 || i%5 == 0 {
		if isMultipleOf3 || isMultipleOf5 {
			//sum = sum + i
			sum += i
		}
	}

	fmt.Println(sum)

	//fmt.Println(time.Since(start))
}
