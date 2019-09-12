# Methods
## Methods (https://tour.golang.org/methods/1)
- Talvez pareça meio estranho mas...Go não tem classes porém você pode declarar métodos.
- Quem pode ter métodos? Types!
- Para saber a diferença entre um method e uma function, repare que method tem uma variável definida logo após a palavra "func" e ela que determina de qual tipo é o method.

```
package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

// Method
func (p Person) Describe() string {
	return fmt.Sprintf("Name: %s Age: %d", p.name, p.age)
}

// Function
func Describe(p Person) string {
    return fmt.Sprintf("Name: %s Age: %d", p.name, p.age)
}

func main() {
	p1 := Person{"Bob", 40}
	fmt.Println(p1.Describe())
    fmt.Println(Describe(p1))
}
```

- Method pode ser definido também para types que não são structs.

```
package main

import (
	"fmt"
)

type MyNumber int

func (m MyNumber) isOdd() bool {
	return m % 2 != 0
}

func main() {
	m1 := MyNumber(3)
	fmt.Println(m1.isOdd())
}
```

## Pointer receivers
- Vocẽ pode fazer um método receber um ponteiro para um tipo ao invés de receber o valor. Isso é útil para alterar a variável em si.
- Vamos dar uma revisada em pointeiros e analisar o exemplo abaixo.

```
package main

import "fmt"

func changeNumber(p *int) {
	*p = 8
}

func dontChangeNumber(p *int) {
	p = new(int)
	*p = 9
}

func main() {
	var i *int
	
	j := 5 // short declaration variable
	
	k := new(int) // returns a pointer to a int type
	
	fmt.Printf("Just declared - i: %+v  j: %+v k: %+v k value: %+v \n", i, j, k, *k)

	changeNumber(&j)
	
	fmt.Printf("After changeNumber(&j) - i: %+v  j: %+v k: %+v k value: %+v \n", i, j, k, *k)

	dontChangeNumber(&j)
	
	fmt.Printf("After dontChangeNumber(&j) - i: %+v  j: %+v k: %+v k value: %+v \n", i, j, k, *k)
	
	// panic: runtime error: invalid memory address or nil pointer dereference
	//changeNumber(i) 
	//fmt.Printf("After changeNumber(i) - i: %+v  j: %+v k: %+v k value: %+v \n", i, j, k, *k)
	
	dontChangeNumber(i)
	
	fmt.Printf("After dontChangeNumber(i) - i: %+v  j: %+v k: %+v k value: %+v \n", i, j, k, *k)
	
	changeNumber(k)
	
	fmt.Printf("After changeNumber(k) - i: %+v  j: %+v k: %+v k value: %+v \n", i, j, k, *k)
	
	dontChangeNumber(k)
	
	fmt.Printf("After dontChangeNumber(k) - i: %+v  j: %+v k: %+v k value: %+v \n", i, j, k, *k)
	
}
```

Saída:

```
Just declared - i: <nil>  j: 5 k: 0x414020 k value: 0 
After changeNumber(&j) - i: <nil>  j: 8 k: 0x414020 k value: 0 
After dontChangeNumber(&j) - i: <nil>  j: 8 k: 0x414020 k value: 0 
After dontChangeNumber(i) - i: <nil>  j: 8 k: 0x414020 k value: 0 
After changeNumber(k) - i: <nil>  j: 8 k: 0x414020 k value: 8 
After dontChangeNumber(k) - i: <nil>  j: 8 k: 0x414020 k value: 8 
```

- Vamos começar com a declaração de variáveis. 

- **var i \*int**: quando declaramos dessa forma, "i" terá o [zero value](https://golang.org/ref/spec#The_zero_value) correspondente ao seu [tipo](https://golang.org/ref/spec#Pointer_types) sendo, no caso, "nil". 

- **j := 5** : essa forma é chamada de ["short declaration"](https://golang.org/ref/spec#Short_variable_declarations). Curioso? Veja a especificação:

[shortVarDecl](https://github.com/golang/go/blob/43e8fd4ef1ae24f1505bd34708fc30aa2b736c52/src/go/parser/parser.go#L141)
```func (p *parser) shortVarDecl(decl *ast.AssignStmt, list []ast.Expr) {```

[NewObj](https://github.com/golang/go/blob/f90e89e675443731e36c2de4bcd3cdd7316d3dfc/src/go/ast/scope.go#L85)
```func NewObj(kind ObjKind, name string) *Object```

- **k := new(int)** : new é uma builtin que retorna um ponteiro do tipo int. Curioso? Veja a especificação:

[new](https://github.com/golang/go/blob/d152ff286f6ef2b25bd95bf97a429a1dc40ba4b5/src/builtin/builtin.go#L191)
```func new(Type) *Type```

[newobject](https://github.com/golang/go/blob/f90e89e675443731e36c2de4bcd3cdd7316d3dfc/src/runtime/malloc.go#L1027)
```func newobject(typ *_type) unsafe.Pointer```

[Pointer](https://github.com/golang/go/blob/f90e89e675443731e36c2de4bcd3cdd7316d3dfc/src/unsafe/unsafe.go#L17) represents a pointer to an arbitrary type. 

- Por isso, temos a saída:
```Just declared - i: <nil>  j: 5 k: 0x414020 k value: 0 ```

- Analisando **changeNumber(&j)**. Estamos passando o endereço da memória da varíavel j, ou seja, &j. Esse endereço é passado por valor e o recebemos na varíavel p. Usando "*p" fazemos com que esse endereço aponte para o valor "8".

- Analisando **dontChangeNumber(&j)**. Estamos passando o endereço da memória da varíavel j, ou seja, &j. Esse endereço é passado por valor e o recebemos na varíavel p. Mudamos o valor da varíavel p em "p = new(int)" e no endereço recebido da builtin new, apontamos para o valor "9". Note que isso não alterou o valor de j pois conforme especificação de chamada de função em Go, passamos o endereço da memória da varíavel j POR VALOR. Assim, "p = new(int)" não altera nossa varíavel.
```
After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution.
```

Referências: 

https://golang.org/ref/spec#Calls

https://medium.com/@kavehmz/call-by-value-in-go-even-for-references-2bc4b2a3c590

Muito louco, né?

- Analisando **changeNumber(i)**. Aqui ocorre erro pois tentamos apontar "nil" para o valor "8".

- Analisando **dontChangeNumber(i)**. Aqui passamos o valor "nil" mas ele não é utilizado. Não altera nada.

- Analisando **changeNumber(k)**. Note que aqui passamos k direto por já ser um endereço de memória, sem precisar usar o "&'. O valor é alterado e aponta para "8".

- Analisando **dontChangeNumber(k)**. Ocorre o mesmo que já explicamos antes. A varíavel k não sofre alteração.

- ATENÇÃO PARA A PEGADINHA
- Veja o programa abaixo e deduza qual será a saída
```
methods with pointer receivers take either a value or a pointer as the receiver when
```
- Resposta:

# Exercícios
1) Crie um método que calcula quantos anos a pessoa têm a partir da sua data de nascimento (não precisa se preocupar muito com precisão).

Dica: https://golang.org/pkg/time/#Time.Sub
Tema: Methods

2) Aproveitando a struct Person do primeiro exemplo, crie um método que altere o atributo name.
Tema: Pointer receivers

