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

- Se não precisar de algum dos valores, basta substituir por "_".
```
for _, v := range a {
```

## Maps (https://tour.golang.org/moretypes/19)
- Um map é um grupo de elementos de mesmo tipo indexados por um conjunto de chaves distintas de determinado tipo. Por exemplo:
```
var m map[int]string
```
É um conjunto de strings identificados por um int.
- Da forma como foi declarada a variável m, trata-se de um map não inicializado logo seu valor é nil e não é possível adicionar itens nele.
- Para inicializar um map, utiliza-se o "make". Por exemplo:
```
m = make(map[int]string)
```
- Se o map não for inicializado, ocorrerá o erro:
```
panic: assignment to entry in nil map
```
- Há questionamentos de por qual motivo ao se criar um map, já não é inicializado (veja mais em [Autovivification](https://en.m.wikipedia.org/wiki/Autovivification)). Um map é um ponteiro para uma hash table. Logo, se ela não existe, seu valor só pode ser nil e daí a necessidade de inicializar a variável com o make que criará a hash table a ser utilizada.
- Referências: 
https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
https://golang.org/ref/spec#Map_types
https://github.com/golang/go/blob/master/src/runtime/map.go
- É possível também declarar um map já com valores (literal)
```
package main

import "fmt"

var m = map[int]string{0: "abc"}

func main() {
	fmt.Println(m)
}
```
Saída:
```
map[0:abc]
```
- Para inserir um elemento no map, basta determinar chave e valor:
```
m[0] = "abc"
```
- Para utilizar um item:
```
fmt.Println(m[0])
```
- Para deletar um item, utiliza-se o comando "delete":
```
delete(m, 0)
```
- Se precisar verificar se um item existe ou não:
```
v, ok := m[0]
```
Se existir, a variável "ok" estará setada como "true"

## Function values / Function closures (https://tour.golang.org/moretypes/24)
- Você pode passar funções como valores também.
```
package main

import (
	"fmt"
)

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {
	sum := func(x, y int) int {
		return x + y
	}
	
	multiply := func(x, y int) int {
		return x * y
	}
	
	fmt.Println(compute(sum))
	fmt.Println(compute(multiply))
}
```
- Closure é uma função passada por valor que consegue lidar com variáveis externas (fora do seu 'body').
Referência: https://gobyexample.com/closures

```
package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()
	
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}
```

Saída:

```
1
2
3
1
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

7) Faça o exercício proposto em:
https://tour.golang.org/moretypes/18
Tema: Slices

8) Faça o exercício proposto em:
https://tour.golang.org/moretypes/23
Tema: Maps

9) Os programas abaixo vão imprimir a mesma coisa?
```
package main

import "fmt"

func Foo(a []int) {
    a[0] = 8
}

func main() {
    a := []int{1, 2}
    Foo(a)         
    fmt.Println(a) 
}
```

``` 
package main

import "fmt"

func Foo(a [2]int) {
    a[0] = 8
}

func main() {
    a := [2]int{1, 2}
    Foo(a)         
    fmt.Println(a) 
}
```
Temas: Arrays e Slices
Referência (não ver antes de responder): https://yourbasic.org/golang/

10) Responda a questão disponível em https://yourbasic.org/golang/gotcha-copy-missing/

"Why does the copy disappear?"
```
var src, dst []int
src = []int{1, 2, 3}
copy(dst, src) // Copy elements to dst from src.
fmt.Println("dst:", dst)
```

Saída:
```
dst: []
```

11) [OPCIONAL] [PEGADINHA] Responda a questão disponível em https://yourbasic.org/golang/gotcha-range-copy-array/

"Why doesn’t the iteration variable x notice that a[1] has been updated?"
```
var a [2]int
for _, x := range a {
    fmt.Println("x =", x)
    a[1] = 8
}
fmt.Println("a =", a)
```

Saída:
```
x = 0
x = 0        <- Why isn't this 8?
a = [0 8]
```
Tema: Arrays

12) Faça o exercício proposto em: https://tour.golang.org/moretypes/26
Tema: Function closure
