package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

/*
	Em Go, uma struct é um tipo de dados que permite agrupar diferentes tipos de dados em uma única unidade. Ela é semelhante a uma classe em outras linguagens de programação orientada a objetos, mas sem a capacidade de definir métodos.

Uma struct é definida por uma coleção de campos, cada um com seu próprio tipo e nome. Os campos de uma struct são acessados usando a notação de ponto (.).
*/
