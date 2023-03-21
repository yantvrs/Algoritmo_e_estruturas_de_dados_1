package main

import "fmt"

type ArrayList struct {
	items []int
	cap   int
}

// Inicia a execução de inicialização
func (list *ArrayList) Init() {
	list.items = make([]int, 10)
}

// Dobra o tamanho do arrayList
func (list *ArrayList) double() {

}

// Adiciona um novo elemento à lista
func (list *ArrayList) Add(item int) {
	list.items = append(list.items, item)
}

// Remove um elemento da lista pelo índice
func (list *ArrayList) Remove(index int) {
	list.items = append(list.items[:index], list.items[index+1:]...)
}

// Retorna o número de elementos na lista
func (list *ArrayList) Size() int {
	return len(list.items)
}

// Retorna o elemento na posição especificada na lista
func (list *ArrayList) Get(index int) int {
	return list.items[index]
}

func main() {
	// Cria uma nova lista vazia
	list := ArrayList{}

	// Adiciona alguns elementos à lista
	list.Add(10)
	list.Add(20)
	list.Add(30)

	// Imprime o tamanho da lista e seus elementos
	fmt.Println("Tamanho da lista:", list.Size())
	fmt.Println("Elementos da lista:", list.items)

	// Remove o segundo elemento da lista
	list.Remove(1)

	// Imprime o tamanho da lista e seus elementos novamente
	fmt.Println("Tamanho da lista após a remoção:", list.Size())
	fmt.Println("Elementos da lista após a remoção:", list.items)

	// Obtém e imprime o primeiro elemento da lista
	first := list.Get(0)
	fmt.Println("Primeiro elemento da lista:", first)
}
