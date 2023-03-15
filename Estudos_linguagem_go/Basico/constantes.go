package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "ljfjflj"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/*
	Constantes
Constantes são declaradas como variáveis, mas com a palavra-chave const.

Constantes podem ser seqüências de caracteres, booleanos, ou valores numéricos.

Constantes não podem ser declaradas usando a sintaxe :=.
*/
