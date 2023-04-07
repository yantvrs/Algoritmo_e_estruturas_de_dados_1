package main

import "fmt"

type ArrayList struct {
	items []int // Array que armazena os elementos da lista
	size  int   // Quantidade de elementos da lista
}

// Inicialização da lista
func (list *ArrayList) Init() {
	list.items = make([]int, 10)
}

// Duplicação da capacidade
func (list *ArrayList) Double() {
	doubleSize := make([]int, len(list.items)*2)

	for i := 0; i < len(list.items); i++ {
		doubleSize[i] = list.items[i]
	}

	list.items = doubleSize
}

// Adicionar elemento no final
func (list *ArrayList) AddEnd(newItem int) {

	// Dobra o tamanho da lista caso seja necessário
	if len(list.items) == list.size {
		list.Double()
	}

	// Adiciona o novo item no final da lista
	list.items[list.size] = newItem
	list.size++
}

// Adicionar elemento em uma posição específica
func (list *ArrayList) AddIndex(newItem int, index int) {

	// Dobra o tamanho da lista caso seja necessário
	if len(list.items) == list.size {
		list.Double()
	}

	// Realocando valores
	for i := list.size; i > index; i-- {
		list.items[i] = list.items[i-1]
	}

	// Adiciona novo item na posição especificada
	list.items[index] = newItem
	list.size++
}

// Remover elemento no final
func (list *ArrayList) RemoveEnd() {
	if list.size > 0 {
		list.size--
	}

}

// Remover elemento em uma posição específica
func (list *ArrayList) RemoveIndex(index int) {
	if index > 0 && index < list.size {
		for i := index; i < list.size; i++ {
			list.items[i] = list.items[i+1]
		}
		list.size--
	}
}

// Obter elemento em uma dada posição da lista
func (list *ArrayList) Get(index int) int {
	if index >= 0 && index < list.size {
		return list.items[index]
	}
	return -1
}

// Trocar um elemento por outro
func (list *ArrayList) Set(newValue int, index int) {
	if index >= 0 && index < list.size {
		list.items[index] = newValue
	}
}

// Tamanho da lista
func (list *ArrayList) Size() int {
	return list.size
}

// Função principal
func main() {
	list := ArrayList{}
	list.Init()

	list.AddEnd(1)
	list.AddEnd(2)
	list.AddEnd(3)
	list.AddEnd(4)
	list.AddEnd(5)
	list.AddEnd(6)
	list.AddEnd(7)
	list.AddEnd(8)
	list.AddEnd(9)
	list.AddEnd(10)

	for i := 0; i < list.Size(); i++ {
		fmt.Print(list.Get(i), " ")
	}
	fmt.Println()

	list.AddIndex(15, 3)

	for i := 0; i < list.Size(); i++ {
		fmt.Print(list.Get(i), " ")
	}
	fmt.Println()

	list.RemoveEnd()
	for i := 0; i < list.Size(); i++ {
		fmt.Print(list.Get(i), " ")
	}
	fmt.Println()

	list.RemoveIndex(2)
	for i := 0; i < list.Size(); i++ {
		fmt.Print(list.Get(i), " ")
	}
	fmt.Println()

	list.Set(25, 5)
	for i := 0; i < list.Size(); i++ {
		fmt.Print(list.Get(i), " ")
	}
	fmt.Println()

	fmt.Println("Tamanho da lista após as alterações:", list.Size())

}
