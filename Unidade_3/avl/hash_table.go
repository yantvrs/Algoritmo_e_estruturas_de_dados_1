package main

import("fmt" "errors")

type Map interface {
	Put(key int, value string)
	Get(key int) (string, error)
	Remove(key int) error
	Size() int
	LoadFactor() float32
	Init()
}

/* Usando generics
type Pair[T, U any] struct {
    key  T
    value U
}*/

/*
	A estrutura `Tuple` é definida como uma representação de um par chave-valor na tabela hash. Ela possui dois campos:

1. O campo `key` é do tipo `int` e representa a chave associada ao valor.
2. O campo `value` é do tipo `string` e representa o valor associado à chave.

Essa estrutura é utilizada na implementação da tabela hash para armazenar os elementos inseridos, onde cada elemento é representado por um objeto `Tuple`.

Ao utilizar a estrutura `Tuple`, é possível associar um valor (representado pelo campo `value`) a uma chave (representada pelo campo `key`). Essa associação é importante para permitir a recuperação eficiente de valores a partir de suas chaves correspondentes na tabela hash.

Por exemplo, se tivermos um objeto `Tuple` com `key` igual a 5 e `value` igual a "Hello", significa que a chave 5 está associada ao valor "Hello". Essa associação é feita para cada elemento inserido na tabela hash.

A estrutura `Tuple` é uma maneira de encapsular e agrupar informações relacionadas (chave e valor) em um único objeto, tornando mais fácil manipular e trabalhar com pares chave-valor na tabela hash.
*/
type Tuple struct {
	key   int
	value string
}
/*
A linha de código `var emptyTuple Tuple = Tuple{}` cria uma variável chamada `emptyTuple` que é do tipo `Tuple` e atribui a ela um valor inicial vazio.

A expressão `Tuple{}` cria uma nova instância da estrutura `Tuple` com seus campos `key` e `value` definidos com valores padrão. No caso de `int`, o valor padrão é `0`, e para `string`, o valor padrão é uma string vazia `""`.

Assim, a variável `emptyTuple` é inicializada como uma instância vazia da estrutura `Tuple`, onde seus campos `key` e `value` possuem os valores padrão `0` e `""`, respectivamente.

Essa variável `emptyTuple` pode ser utilizada, por exemplo, como um marcador ou indicador para representar um espaço vazio em uma tabela hash ou indicar a ausência de um elemento associado a uma chave específica.
*/
var emptyTuple Tuple = Tuple{}

type HashTable struct {
	buckets          []Tuple
	elementsInserted int
}

