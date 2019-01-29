# Flow Control
## For (https://tour.golang.org/flowcontrol/1)
- Ele é composto por: init (executado antes da 1 iteração), condição (verificada antes de cada iteração) e pós (executado no final de cada iteração)
- Exemplo: 
```
for i:= 0; i < 10; i++
```
- Mas e o while? Desapega!
- Exemplo:
```
for sum < 10 { 
    sum += sum
} 
```
## If (https://tour.golang.org/flowcontrol/5)
- O if é composto por uma expressão a ser validada. Sendo true, o que estiver entre {} será executado
- Além disso, você pode executar algo antes de validar a condição lembrando que qualquer declaração de variável será válida apenas DENTRO do escopo do if
- Exemplo
```
limit := 10
if sum := 4 + 4; sum < limit {
	fmt.Println(sum)
}
```

## Switch (https://tour.golang.org/flowcontrol/9)
- A diferença do switch do Go para o switch de outras linguagens é que ele já inclui o break automaticamente no final de cada "case"
- Exemplo
```
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Should I go to work today?")
	
	weekday := time.Now().Weekday().String();
	
	switch weekday {
		case "Saturday":
			fmt.Println("No.")
		case "Sunday":
			fmt.Println("No.")
		default:
			fmt.Println("Yes.")
	}
}
```
- Para cair no próximo case (mesmo se a expressão não for válida), pode-se usar o "fallthrough".

## Defer (https://tour.golang.org/flowcontrol/12)
- O defer indica que uma função/comando deve ser executada apenas no fim da função em que está contextualizado. Similar ao "finally" do Java.
- Exemplo
```
package main

import "fmt"

func main() {
	defer fmt.Println("Bye.")

	fmt.Println("Hello.")
}
```
- A função dentro do defer pode ler e alterar o que a função que contêm o defer retorna
- Exemplo
```
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

# Exercícios
1) Quantas vezes a palavra "banana" será impressa?
```
package main

import "fmt"

func main() {
	sum := 1
	for  ; sum < 8; {
		sum += sum
		fmt.Println("banana")
	}
}
``` 

1.1) E agora?
```
package main

import "fmt"

func main() {
	sum := 1
	for  sum = 2; sum < 8; {
		sum += sum
		fmt.Println("banana")
	}
}
```

1.2) Esse programa vai imprimir o mesmo que o do exercício anterior?
```
package main

import "fmt"

func main() {
        for  sum = 2; sum < 8; {
                sum += sum
                fmt.Println("banana")
        }
}
```

Tema: For

2) Faça um programa que imprima a palavra "Golang" infinitamente.
Tema: For

3) O que o programa abaixo irá imprimir?
```
package main

import (
	"fmt"
	"math"
)

func main() {
        square := 0
	if square := math.Sqrt(9); square < 10 {
		fmt.Println(square)
	}
	fmt.Println(square)
}
```
Tema: If

4) Adapte o exemplo do item "switch" para que sábado e domingo sejam tratados de forma igual (dica: fallthrough).
Tema: Switch

5) O que a função abaixo irá imprimir?
```
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```
Tema: Defer
