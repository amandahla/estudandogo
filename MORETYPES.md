# More Types
## Pointers (https://tour.golang.org/moretypes/1)
- Não há como escapar deles...Go também tem ponteiros!
```
var p *int
```
- A variável p é um ponteiro para um valor do tipo int. Seu valor zero é nil.
- Para obter o ponteiro de uma variável, utiliza-se o operador &
```
sum := 50
p = &sum
```
- A variável p é um ponteiro para a variável sum
- Para obter o valor de uma variável por seu ponteiro, utiliza-se o operador *
- O ato de obter um valor por ponteiro é chamado de "dereference" ou "indirecting"
```
fmt.Println(*p)
```

## Structs (https://tour.golang.org/moretypes/2)
- "A struct is a sequence of named elements, called fields, each of which has a name and a type." ([fonte](https://golang.org/ref/spec#Struct_types))
- Alguns comparam a ideia de "struct" com a de "classe" no contexto de Orientação a Objeto.
```
type Person struct {
	Name string
	Age int
}

p1 := Person{Name: "Clara", Age: 16}

p2 := Person{"Bruno", 14}
```
- Fique ligado: como vimos em Packages, variáveis cujo nome começa com letra maiúscula são exportadas!
- Para acessar um campo, usa-se variavelStruct.NomedoCampo
```
p3 := Person{Name: "Roberto", Age: 35}
fmt.Println(p3.Name)
```
- Isso vale também para acessar o campo de um ponteiro de struct. Como a notação ``(*p).Field`` ficaria pesada, Go permite escrever como p.Field
```
k := &p3
fmt.Println(k.Name)
```

## Arrays (https://tour.golang.org/moretypes/6)
- Um Array é uma sequência de elementos de mesmo tipo. 
- A quantidade de elementos faz parte do tipo Array e é definida na declaração da variável. Assim, um array não pode ser redimensionado.
```
var numbers [10]int
```

# Exercícios
1) O que o programa abaixo irá imprimir?
```
package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 15
	*b = 14
	
	fmt.Println(a)
}
``` 
Tema: Pointers

2) Sem trapacear: Só de bater o olho, qual o erro do programa abaixo?
```
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	p1 := Person{"Caio"}
	fmt.Println(p1.Name)
}
```
Tema: Structs

3) O programa abaixo está correto?
```
package main

func main() {
	var numbers [2]int
	numbers[0] = 3
	numbers[1] = 2
	numbers[2] = 1
	numbers[3] = 0
}
```
