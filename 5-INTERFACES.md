# Methods
## Interfaces (https://tour.golang.org/methods/9)
- o tipo Interface é bem similar ao conceito no Java para criar um conjunto de assinaturas de métodos a serem implementados por outros tipos. Vamos ver um exemplo para entender melhor:
```
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

func (d Dog) Walk() {
	fmt.Println("walking")
}

func (d Dog) Bark() {
 	fmt.Println("barking")
}

func DoSomething(a Animal) {
	a.Walk()
}

func main() {
	thor := Dog{"Thor"}
	DoSomething(thor)
	thor.Bark()
}
```

- Aqui Dog é do tipo Animal pois implementa o método Walk logo pode ser enviado como parâmetro para a função DoSomething
- Note que assim interfaces são implementadas de forma **implícita**, ou seja, não há no código alguma indicação que o tipo Dog implementa a interface Animal.
- E se mudarmos o método "func (d Dog) Walk()" para "func (d *Dog) Walk()"? Vou estragar a surpresa. Irá ocorrer o erro "cannot use thor (type Dog) as type Animal in argument to DoSomething: Dog does not implement Animal (Walk method has pointer receiver)"
- Isso porque quando criamos a variável "thor" ela é do tipo Dog que não é o tipo esperado do método Walk. Para consertar, vamos então mudar de "thor := Dog{"Thor"}" para "thor := &Dog{"Thor"}"
- É importante ficar atento aqui que se estamos trabalhando com interfaces, pode acontecer de recebermos um "nil" no nosso método "Walk()".
- Há ainda a interface vazia ("empty interface") que não especifica nenhum método e é utilizada para lidar com valores de tipos desconhecidos. Por exemplo, o fmt.Print aceita parâmetros do tipo "interface{}". Veja o exemplo no link do [Tour of Go](https://tour.golang.org/methods/14).


# Exercícios
4) Experimente remover o método "func (d Dog) Walk()" para observar o erro que ocorre.
5) Modifique o método "Walk()" para não ocorrer erro se receber um "nil".
