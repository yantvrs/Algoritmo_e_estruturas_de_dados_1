package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)
}

/*
	O exemplo mostra vários tipos de variáveis e também que as declarações de variáveis podem ser "construídas" em blocos, como com as declarações de importação.

Os tipos int, uint e uintptr são geralmente de 32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits. Quando você precisar de um valor inteiro deverá usar int, a menos que tenha um motivo específico para usar um tipo de inteiro com tamanho especificado ou sem sinal.




	Aqui estão os principais tipos de dados em Go:

Tipos básicos:
int: inteiros com sinal, como int8, int16, int32 e int64.
uint: inteiros sem sinal, como uint8, uint16, uint32 e uint64.
float: números de ponto flutuante, como float32 e float64.
complex: números complexos, como complex64 e complex128.
bool: valores booleanos verdadeiro ou falso.
string: sequências de caracteres.

Tipos compostos:
array: uma coleção fixa de elementos do mesmo tipo.
slice: uma coleção dinâmica de elementos do mesmo tipo.
map: uma coleção de pares chave-valor não ordenados.
struct: uma coleção de campos nomeados que podem ter diferentes tipos.
interface: um tipo que define um conjunto de métodos que uma estrutura pode implementar.

Além disso, Go também possui tipos especiais, como:

pointer: um tipo que armazena o endereço de memória de uma variável.
function: um tipo que define uma função que pode ser passada como argumento ou retornada por outra função.
channel: um tipo usado para comunicação entre goroutines, as unidades concorrentes de execução em Go.
*/
