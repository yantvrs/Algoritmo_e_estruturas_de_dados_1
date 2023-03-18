package main // Declara o nome do pacote como main

import (
	"fmt"  // Importa o pacote fmt para usar funções de entrada e saída
	"math" // Importa o pacote math para usar funções matemáticas
)

type Vertex struct { // Define um novo tipo Vertex como uma estrutura com dois campos
	X, Y float64 // Os campos X e Y são números de ponto flutuante
}

func Abs(v Vertex) float64 { // Define uma função Abs que recebe um valor Vertex como parâmetro e retorna um float64
	return math.Sqrt(v.X*v.X + v.Y*v.Y) // Calcula a raiz quadrada da soma dos quadrados das coordenadas do vértice usando a função Sqrt do pacote math
}

func Scale(v *Vertex, f float64) { // Define uma função Scale que recebe um ponteiro para Vertex e um float64 como parâmetros
	v.X = v.X * f // Multiplica o campo X do vértice pelo parâmetro f
	v.Y = v.Y * f // Multiplica o campo Y do vértice pelo parâmetro f
}

func main() { // Define a função principal que é executada quando o programa é iniciado
	v := Vertex{3, 4}   // Cria uma variável v do tipo Vertex e inicializa seus campos com 3 e 4
	Scale(&v, 10)       // Chama a função Scale com o endereço de v e 10 como parâmetros, alterando os valores de X e Y de v para 30 e 40 respectivamente
	fmt.Println(Abs(v)) // Chama a função Abs com v como parâmetro e imprime seu resultado usando a função Println do pacote fmt. O resultado é aproximadamente 50.
}

/*

Ponteiros e funções
Aqui vemos os métodos da Scale e Abs reescritos como funções.

Mais uma vez, tente remover o * a partir da linha 16. Você pode ver porque as mudanças de comportamento ocorrem? O que mais você precisa mudar para o exemplo compilar?

(Se você não tiver certeza, continue para a próxima página.)
*/
