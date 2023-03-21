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
	doubleItems := make([]int, len(list.items)*2)

	for i := 0; i < len(list.items); i++ {
		doubleItems[i] = list.items[i]
	}

	list.items = doubleItems
}

// Adiciona um novo elemento à lista
func (list *ArrayList) Add(newItem int) {

	// Dobra a arrayList caso não haja espaço suficiente
	if list.cap == len(list.items) {
		list.double()
	}

	list.items[list.cap] = newItem
	list.cap++
}

// Adiciona um novo elemento em uma posição específica
func (list *ArrayList) AddOnIndex(newItem int, index int) {

	// Dobra a arrayList caso não haja espaço suficiente
	if list.cap == len(list.items) {
		list.double()
	}

	// Alocando os items em novas posições
	for i := list.cap; i > index; i-- {
		list.items[i] = list.items[i-1]
	}

	list.items[index] = newItem
	list.cap++
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
