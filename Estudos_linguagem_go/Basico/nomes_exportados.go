package main

import (
	"fmt"
)

func main() {
	fmt.Println(math.pi)
}

/*
	Em Go, um nome é exportado se ele começa com uma letra maiúscula. Por exemplo, Pizza é um nome exportado, assim como Pi, que é exportado do pacote math.

pizza e pi não começam com uma letra maiúscula, logo eles não são exportados.

Ao importar um pacote, você pode referenciar apenas seus nomes exportados. Todos os nomes "não exportados" não são acessíveis de fora do pacote.

Execute o código. Observe a mensagem de erro.

Para corrigir o erro, renomeie math.pi para math.Pi e tente novamente.


Em Go, existem diferentes tipos de funções para imprimir conteúdo no console ou em outros dispositivos de saída, cada uma com sua própria sintaxe e uso adequado. Aqui estão alguns dos principais:

fmt.Print: Imprime uma string simples sem formatação adicional. É o mais básico dos métodos de impressão. Exemplo:

python
Copy code
fmt.Print("Olá, mundo!")
fmt.Println: Imprime uma string com uma nova linha adicionada no final. É muito útil para separar as saídas em linhas diferentes. Exemplo:

python
Copy code
fmt.Println("Olá,")
fmt.Println("mundo!")
fmt.Printf: Imprime uma string formatada, onde pode-se incluir placeholders para variáveis ou expressões. É semelhante à função printf em C. Exemplo:

go
Copy code
idade := 25
nome := "João"
fmt.Printf("%s tem %d anos.\n", nome, idade)
fmt.Sprintf: Funciona de maneira semelhante ao fmt.Printf, mas em vez de imprimir na saída padrão, retorna a string formatada como resultado. Exemplo:

go
Copy code
idade := 25
nome := "João"
resultado := fmt.Sprintf("%s tem %d anos.", nome, idade)
fmt.Println(resultado)
log.Print, log.Println e log.Printf: Essas funções estão no pacote log, que é útil para imprimir mensagens de log e depuração. Elas são semelhantes aos métodos de impressão do pacote fmt, mas adicionam informações adicionais, como a hora em que a mensagem foi registrada. Exemplo:

python
Copy code
import "log"
log.Println("Erro: Falha ao carregar arquivo.")
Esses são apenas alguns dos métodos de impressão disponíveis em Go. Cada um tem seus próprios recursos e usos adequados, e pode ser útil para diferentes cenários e necessidades de impressão.
*/
