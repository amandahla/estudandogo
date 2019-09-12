# Basics
## Packages (https://tour.golang.org/basics/1)
- Todo programa Go inicia no package "main"
- Para usar outros packages, é preciso listá-los no "import"
- O nome do package a ser utilizado no programa é a última string no "import". Por exemplo, "math/rand" é chamado de "rand"

## Exported names (https://tour.golang.org/basics/3)
- Quando você importa um package, só poderá "chamar" o que estiver exportado.
- Em Go, para exportar algo, é preciso iniciar seu nome com letra maiúscula. Por exemplo, usar"math.pi" irá gerar erro enquanto "math.Pi" terá o comportamento esperado.

## Functions (https://tour.golang.org/basics/4)
- Para declarar uma função, usa-se:
func nomeDaFuncao(nomeDoParametro tipoDoParametro) tipoDoRetorno
Exemplo: 
```
func add(x int, y int) int{
	return x + y
}
```

Outra forma utilizada é determinar nome para os valores a serem retornados
Exemplo: 
```
func add(x,y int) (result int) {
	result = x + y
	return
}
```
O ideal é usar esse recurso apenas em funções pequenas para não prejudicar a legibilidade do programa.

- Além disso, é possível retornar múltiplos valores. Por exemplo:
```
func addMult(x,y int) (add,mul int) {
	add = x + y
	mul = x * y
	return
}
```

## Variables (https://tour.golang.org/basics/8)
- É possível declarar variáveis tanto no contexto package quanto function. Assim, uma variável declarada fora de func, vale para todo o programa.
- Lembrando: var nomeVariavel tipoVariavel
- Se for inicializar: var nomeVariavel tipoVariavel = valor ou var nomeVariavel = valor (o tipo será inferido)
- Só é possível usar o ":=" ao invés do "var" dentro do contexto function.
- Para converter uma variável de um tipo para outro, pode usar tipo(nomeVariavel). Por exemplo:
```
var i int = 42
var f float64 = float64(i)
```

##Constants (https://tour.golang.org/basics/15)
- São declaradas como variáveis mas ao invés de "var", utiliza-se "const"
- Não se pode declarar com ":="
- Exemplo:
``const mensagem = "olá"``
- A diferença entre "var" e "const" é que não é permitido alterar o valor de uma const. Irá ocorrer o erro:
``cannot assign to nomeVariavel``

# Exercícios
1) Resolver o problema 1 do Projeto Euler.
Se listamos todos os números naturais inferiores a 10 que são múltimos de 3 ou 5, temos 3,5,6,e 9. A soma desses valores é 23. O programa deve imprimir na tela a soma dos múltiplos de 3 ou 5 inferiores a 1000.
Temas: Packages, Variables

2) Altere o programa criado anteriormente para que o número máximo seja aleatório. 
Dica: "math/rand" e não esqueça de usar "rand.Seed(time.Now().UTC().UnixNano())" para não retornar sempre o mesmo valor
Imprima na tela o valor do número máximo com a frase "Maximum number is " e depois a soma.
Tema: Exported names

3) Altere o programa criado anteriormente para que o número máximo seja obtido por meio de uma função getRandomNumber(n) em que n é o valor máximo a ser gerado.
Tema: Functions

4) Altere o programa criado anteriormente para que a função getRandomNumber(n) retorne o número e a mensagem a ser exibida "Maximum number is". 
Dica: message = fmt.Sprintf("Maximum number is %v\n", randomNumber)
Tema: Functions

5) Altere o programa criado anteriormente para:
- Ao invés de exibir somente a soma, deverá exibir "Sum is"
- Tanto para a mensagem da soma quanto para a mensagem retornada pela função getRandomNumber, o programa deverá checar a variável printInPortuguese. Se o valor dela for true, exibirá "Soma é" e "Número máximo é". Se for false, exibirá "Sum is" e "Maximum number is"
Dica: onde a variável printInPortuguese deve ser declarada para ser utilizada tanto pela function main quanto pela getRandomNumber?
Saída de exemplo:
O número máximo é 34
Soma é 258
Tema: Variables

6) Altere o programa criado anteriomente para exibir a raiz quadrada de sum com a mensagem "Sqrt is". Dica: a função Sqrt espera um parâmetro do tipo float64.
Tema: Variables

7) Reflexão: a variável printInPortuguese utilizada nos programas criados anteriormente poderia ser uma constante? Justifique.
Tema: Constants

