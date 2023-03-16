package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

/*
	Ponteiros para structs
Campos de estruturas podem ser acessados através de um ponteiro de estrutura.

Para acessar o campo X de uma struct quando tivermos o ponteiro p da struct podemos escrever (*p).X. No entanto, essa notação é incômoda, então a linguagem nos permite escrever apenas p.X, sem a referência explícita.
*/
