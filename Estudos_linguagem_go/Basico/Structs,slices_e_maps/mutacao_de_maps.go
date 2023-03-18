package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

/*

Mutação de Maps
Inserir ou atualizar um elemento no map m:

m[key] = elem
Recuperar um elemento:

elem = m[key]
Excluir um elemento:

delete(m, key)
Teste que uma chave está presente com dois valores:

elem, ok = m[key]
Se key está em m, ok é true. Se não, ok é false.

Se key não está no map então elem tem valor zero para o elemento do tipo do map.

Nota: Se elem ou ok não foi declarado ainda você pode usar esta forma de declaração curta:

elem, ok := m[key]
*/
