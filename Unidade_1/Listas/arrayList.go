package main

import "fmt"

type ArrayList struct {
	items []int // array que armazena os elementos da lista
	cap   int   // capacidade da lista
}

// Inicializa a lista com uma capacidade inicial de 10 elementos
func (list *ArrayList) Init() {
	list.items = make([]int, 10)
}

// Dobra a capacidade da lista
func (list *ArrayList) double() {
	doubleItems := make([]int, len(list.items)*2)

	for i := 0; i < len(list.items); i++ {
		doubleItems[i] = list.items[i]
	}

	list.items = doubleItems
}

// Adiciona um novo elemento ao final da lista
func (list *ArrayList) Add(newItem int) {

	// Dobra a capacidade da lista caso necessário
	if list.cap == len(list.items) {
		list.double()
	}

	// Adiciona o novo elemento na próxima posição livre
	list.items[list.cap] = newItem
	list.cap++
}

// Adiciona um novo elemento em uma posição específica da lista
func (list *ArrayList) AddOnIndex(newItem int, index int) {

	// Dobra a capacidade da lista caso necessário
	if list.cap == len(list.items) {
		list.double()
	}

	// Realoca os elementos para abrir espaço para o novo elemento
	for i := list.cap; i > index; i-- {
		list.items[i] = list.items[i-1]
	}

	// Adiciona o novo elemento na posição especificada
	list.items[index] = newItem
	list.cap++
}

// Remove o último elemento da lista
func (list *ArrayList) Remove() {
	if list.cap > 0 {
		list.cap--
	}
}

// Remove um elemento em uma posição específica da lista
func (list *ArrayList) RemoveOnIndex(index int) {
	if index >= 0 && index < list.cap {
		// Realoca os elementos para remover o elemento na posição especificada
		for i := index; i < list.cap; i++ {
			list.items[i] = list.items[i+1]
		}
		list.cap--
	}
}

// Retorna o tamanho atual da lista
func (list *ArrayList) Size() int {
	return list.cap
}

// Retorna o elemento em uma posição específica da lista
func (list *ArrayList) Get(index int) int {
	if index >= 0 && index < list.cap {
		return list.items[index]
	}
	return -1 //TODO: adicionar tratamento de erro
}

// Substitui o elemento em uma posição específica da lista pelo valor especificado
func (list *ArrayList) Set(value, index int) {
	if index >= 0 && index < list.cap {
		list.items[index] = value
	}
	//TODO: adicionar tratamento de erro
}

// Função principal
func main() {
	list := ArrayList{}
	list.Init()

	// Adiciona elementos na lista
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	list.Add(7)
	list.Add(8)
	list.Add(9)
	list.Add(10)

	// Adiciona um novo elemento na posição 5
	list.AddOnIndex(15, 5)

	// Imprime o tamanho da lista
	fmt.Println("Tamanho da lista:", list.Size())

	// Imprime todos os elementos da lista
	for i := 0; i < list.Size(); i++ {
		fmt.Println(list.Get(i))
	}

	// Remove o último elemento da lista
	list.Remove()

	// Remove o elemento na posição 3
	list.RemoveOnIndex(3)

	// Altera o valor do elemento na posição 2 para 20
	list.Set(20, 2)

	// Imprime todos os elementos da lista após as alterações
	fmt.Println("Tamanho da lista após as alterações:", list.Size())
	for i := 0; i < list.Size(); i++ {
		fmt.Println(list.Get(i))
	}
}
