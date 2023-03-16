package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*
	O código possui três linhas de importação de pacotes, que são "fmt", "math/rand" e "main". O pacote "fmt" é utilizado para formatação de saída no console, enquanto "math/rand" é utilizado para gerar números aleatórios. O pacote "main" é o pacote principal que contém a função "main" que é executada quando o programa é iniciado.

A função "main" chama a função "Println" do pacote "fmt" para exibir a mensagem "My favorite number is" e o resultado da função "Intn" do pacote "rand", que gera um número aleatório entre 0 e 9. Ao executar o programa, a saída será a mensagem "My favorite number is" seguida de um número aleatório.
*/
