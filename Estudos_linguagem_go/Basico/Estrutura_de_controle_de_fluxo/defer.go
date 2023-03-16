package main

import "fmt"

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}

/*
	Defer
A declaração defer adia a execução de uma função até o final do retorno da função.

Os argumentos das chamadas adiadas são avaliados imediatamente, mas a função chamada não é executada até o retorno da função.
*/
