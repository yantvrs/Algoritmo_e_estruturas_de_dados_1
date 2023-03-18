package main // Declara o nome do pacote como main

import (
	"fmt"  // Importa o pacote fmt para usar funções de entrada e saída
	"math" // Importa o pacote math para usar funções matemáticas
)

type Vertex struct { // Define um novo tipo Vertex como uma estrutura com dois campos
	X, Y float64 // Os campos X e Y são números de ponto flutuante
}

func (v Vertex) Abs() float64 { // Define um método Abs que recebe um valor Vertex como receptor e retorna um float64
	return math.Sqrt(v.X*v.X + v.Y*v.Y) // Calcula a raiz quadrada da soma dos quadrados das coordenadas do vértice usando a função Sqrt do pacote math
}

func (v *Vertex) Scale(f float64) { // Define um método Scale que recebe um ponteiro para Vertex como receptor e um float64 como parâmetro
	v.X = v.X * f // Multiplica o campo X do vértice pelo parâmetro f
	v.Y = v.Y * f // Multiplica o campo Y do vértice pelo parâmetro f
}

func main() { // Define a função principal que é executada quando o programa é iniciado
	v := Vertex{3, 4}    // Cria uma variável v do tipo Vertex e inicializa seus campos com 3 e 4
	v.Scale(10)          // Chama o método Scale com v como receptor e 10 como parâmetro, alterando os valores de X e Y de v para 30 e 40 respectivamente
	fmt.Println(v.Abs()) // Chama o método Abs com v como receptor e imprime seu resultado usando a função Println do pacote fmt. O resultado é aproximadamente 50.
}

/*
	Ponteiro receptor
Você pode declarar métodos com ponteiro receptor.

Isso significa que o tipo de receptor tem a sintaxe literal *t por algum tipo t. (Além disso, t não pode ser ele próprio, um ponteiro como *int.)

Por exemplo, o método Scale aqui é definido no *Vertex.

Métodos com ponteiro receptores pode modificar o valor ao qual o receptor Os pontos (como Scale faz aqui). Desde métodos muitas vezes precisam modificar seu receptor, receptores de ponteiro são mais comuns do que receptores de valor.

Tente remover o * a partir da declaração da função Scale na linha 16 e observar como o comportamento do programa muda.

Com um receptor de valor, o método Scale opera sobre uma cópia do original `Valor Vertex`. (Este é o mesmo comportamento para qualquer outro argumento da função.) O método Scale deve ter um receptor ponteiro para mudar o `valor Vertex` declarados na função main.
*/
