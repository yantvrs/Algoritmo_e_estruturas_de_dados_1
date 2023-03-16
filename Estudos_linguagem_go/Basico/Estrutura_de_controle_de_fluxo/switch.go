package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux. ")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("")
	}
}

/*
	Switch
Uma instrução switch é uma forma mais curta de escrever uma sequência de declarações if - else. Ele executa o primeiro caso cujo valor é igual à expressão de condição.

O switch do Go é como o do C, C ++, Java, JavaScript e PHP, exceto que o Go apenas executa o caso selecionado, não todos os casos que seguem. De fato, a declaração break que é necessária no final de cada caso naquelas linguagens é fornecido automaticamente no Go. Outra diferença importante é que os cases dos switchs do Go não precisam ser constantes, e os valores envolvidos não precisam ser inteiros.
*/
