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

## Slices (https://tour.golang.org/moretypes/7)
- A diferença entre Array e Slice é que Array tem tamanho fixo e Slice não.
- Pode ser declarado como:
-- []T onde T é o tipo dos elementos do slice
-- [minimo:maximo]T onde o máximo é excluído, ou seja, um slice [1:4]T terá elementos de índices 1, 2 e 3. O valor de minimo é 0 e de máximo é o length do array.
- Um slice não armazena dados, ele faz referência a array. Assim, ao criar um slice com conteúdo de um array existente, se esse array for alterado, o slice exibirá a alteração também.
- Um slice tem duas propriedades: length (comprimento) e capacity (capacidade). Length é o número de elementos que um slice possui e capacity é o número de elementos do array para o qual o slice aponta.
```
package main

import "fmt"

func main() {
	fruits := [3]string{"banana", "apple", "cherry"}

	s1 := fruits[1:2]

	fmt.Println(s1)

	fmt.Printf("len %d cap %d \n", s1)

	fruits[1] = "kiwi"

	fmt.Println(s1)

	s1 = fruits[:1]
	
	fmt.Println(s1)

	fmt.Printf("len %d cap %d \n", s1)

	s1 = fruits[1:]

	fmt.Println(s1)

	fmt.Printf("len %d cap %d \n", s1)
}
```
- Outra maneira de declarar um slice é utilizando a função "make". Com ela, é possível determinar length e capacity do slice.
```
a := make([]int, 5) // len(a) = 5 cap(a) = 5
b := make([]int, 5, 10) // len(b) = 5 cap(b) = 10
```
- Resumindo, se um slice é uma caixa de ovos, length é quantos ovos você têm e capacity é quantos cabem na caixa. Quanto maior a capacity, mais memória a linguagem vai alocar para seu programa. O momento que mais precisamos ter cuidado com isso é ao usar a função "append" que permite adicionar elementos em um slice. Vamos ver como se comporta com cada uma das 3 maneiras de se criar um slice.

-- Declarando com var
```
package main

import "fmt"

func main() {
	var a []int

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
```
Saída:
```
len=0 cap=0 []
len=5 cap=8 [0 1 2 3 4]
```

-- Declarando literalmente
```
package main

import "fmt"

func main() {
	a := []int{}

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
```
Saída:
```
len=0 cap=0 []
len=5 cap=8 [0 1 2 3 4]
```

-- Usando o make
```
package main

import "fmt"

func main() {
	var a = make([]int, 5)

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
```
Saída:
```
len=5 cap=5 [0 0 0 0 0]
len=10 cap=12 [0 0 0 0 0 0 1 2 3 4]
```

-- Parece que estamos esbanjando memória aqui, não? 
Como otimizar o exemplo anterior declarando slice usando o make? 
Por default, o make retorna um slice com capacity igual ao length e por isso o append faz com que os elementos sejam adicionados a partir do 5 elemento do slice. 

Uma solução seria declarar com length 0 e capacity 5.
```
package main

import "fmt"

func main() {
	var a = make([]int, 0, 5)

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
```
Saída:
```
len=0 cap=5 []
len=5 cap=5 [0 1 2 3 4]
```

-- Quer saber de onde veio esse capacity = 8 nos dois primeiros exemplos? Leia a explicação no item "Append: An example" do link:
https://blog.golang.org/slices

- Mudando um pouco de assunto, vamos ver agora como iterar um slice. Uma maneira é utilizando o "range". Ele retorna dois valores: o índice e o conteúdo.
```
package main

import "fmt"

func main(){
	a := []string{"one","two","three"}

	for k, v := range a {
		fmt.Printf("k %d v %s \n", k, v)
	}
}
```

- Se não precisar utilizar algum dos valores, basta substituir por "_".
```
for _, v := range a {
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

4) O programa abaixo está correto?
Tema: Slices
Referências (ver só depois de responder!): 
https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
https://golang.org/ref/spec#For_statements
```
package main

import "fmt"

var t []string

for _, v := range t { 
    	fmt.Println("%v", v)
}

```

5) [OPCIONAL] Compare declaração/iteração de slice (conceito equivalente) com outras linguagens de programação que você conhece. O programa abaixo em Java, por exemplo. O que aconteceria? (arquivo disponível em Exercicios/MORETYPES/5)
```
public class Brincadeira {
	public static void main(String[] args) {
		int ar[];

		for (int i : ar ) {
			System.out.println(i);
		}
	}	
}
```

6) Que outra solução poderia ser implementada para otimizar uso de memória no código abaixo?
Tema: Slices
```
package main

import "fmt"

func main() {
	var a = make([]int, 5)

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}

	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
```
