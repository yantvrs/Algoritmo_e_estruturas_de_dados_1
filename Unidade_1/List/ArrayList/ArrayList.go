package main

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
func (list *ArrayList) AddEnd(newItem int, index int) {

	// Dobra o tamanho da lista caso seja necessário
	if len(list.items) == list.size {
		list.Double()
	}

	// Adiciona o novo item no final da lista
	list.items[list.size] = newItem
	list.size++
}

// Adicionar elemento em uma posição específica
func (list *ArrayList) AddIndex(newItem int) {

	// Dobra o tamanho da lista caso seja necessário
	if len(list.items) == list.size {
		list.Double()
	}

	// Reorganizando os valores

}

// Remover elemento no final

// Remover elemento em uma posição específica

// Obter elemento em uma dada posição da lista

// Trocar um elemento por outro

// Tamanho da lista

// Função principal
