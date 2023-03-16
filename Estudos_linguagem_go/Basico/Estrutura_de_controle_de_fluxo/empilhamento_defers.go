package main

import "fmt"

func main() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

/*

Empilhando defers
Chamadas de funções adiadas são empurradas por uma pilha. Quando a função retorna, as chamadas de adiamento são executadas na ordem em que a última a entrar é a primeira a sair.
*/
