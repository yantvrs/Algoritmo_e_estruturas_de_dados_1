package main

import "fmt"

// Definindo a estrutura do nó da lista encadeada
type Node struct {
	data int   // Valor do nó
	next *Node // Próximo nó da lista
}

// Definindo a estrutura da lista encadeada
type LinkedList struct {
	head *Node // Nó inicial da lista
}

// Função para adicionar um nó no fim da lista
func (list *LinkedList) Add(data int) {
	newNode := &Node{data, nil} // Criando um novo nó

	if list.head == nil { // Se a lista estiver vazia, o novo nó se torna o primeiro
		list.head = newNode
	} else {
		current := list.head // Começando a busca pelo fim da lista

		for current.next != nil { // Percorrendo a lista até encontrar o último nó
			current = current.next
		}

		current.next = newNode // Adicionando o novo nó no fim da lista
	}
}

// Função para remover um nó da lista
func (list *LinkedList) Remove(data int) {
	if list.head == nil { // Se a lista estiver vazia, não há nada para remover
		return
	}

	if list.head.data == data { // Se o nó a ser removido for o primeiro da lista
		list.head = list.head.next // O segundo nó se torna o primeiro
		return
	}

	current := list.head // Começando a busca pelo nó a ser removido

	for current.next != nil {
		if current.next.data == data { // Se o próximo nó for o nó a ser removido
			current.next = current.next.next // Removendo o nó da lista
			return
		}
		current = current.next // Continuando a busca
	}
}

// Função para obter o tamanho da lista
func (list *LinkedList) Size() int {
	size := 0 // Começando com tamanho 0

	if list.head == nil { // Se a lista estiver vazia, retorna tamanho 0
		return size
	}

	current := list.head // Começando a busca pelo fim da lista
	size++               // Adicionando o primeiro nó

	for current.next != nil { // Percorrendo a lista até encontrar o último nó
		size++
		current = current.next
	}

	return size
}

// Função para obter um nó a partir de um índice
func (list *LinkedList) Get(index int) *Node {
	if list.head == nil { // Se a lista estiver vazia, retorna nulo
		return nil
	}

	current := list.head // Começando a busca pelo nó com o índice desejado
	i := 0

	for i < index && current.next != nil { // Percorrendo a lista até encontrar o nó com o índice desejado
		i++
		current = current.next
	}

	if i == index { // Se encontrou o nó com o índice desejado, retorna o nó
		return current
	}

	return nil // Caso contrário, retorna nulo
}

// Função principal
func main() {
	// Criando uma lista encadeada vazia
	list := &LinkedList{}

	// Adicionando alguns nós na lista
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)

	// Imprimindo o tamanho da lista
	fmt.Printf("Tamanho da lista: %d\n", list.Size())

	// Imprimindo o valor do segundo nó
	secondNode := list.Get(1)
	if secondNode != nil {
		fmt.Printf("Valor do segundo nó: %d\n", secondNode.data)
	}

	// Removendo o terceiro nó da lista
	list.Remove(30)

	// Imprimindo o tamanho da lista novamente
	fmt.Printf("Tamanho da lista após remover um nó: %d\n", list.Size())
}
