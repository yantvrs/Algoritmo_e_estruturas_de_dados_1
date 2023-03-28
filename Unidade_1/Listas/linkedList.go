package main

import "fmt"

// Node representa um nó da lista encadeada
type Node struct {
	value int
	next  *Node
}

// LinkedList representa uma lista encadeada
type LinkedList struct {
	head *Node
	size int
}

// Add adiciona um novo nó contendo o valor dado no final da lista encadeada
func (list *LinkedList) Add(value int) {
	newNode := &Node{value, nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}

	list.size++
}

// AddOnIndex adiciona um novo nó contendo o valor dado na posição index da lista encadeada
func (list *LinkedList) AddOnIndex(index int, value int) {
	if index < 0 || index > list.size {
		panic("Index out of bounds")
	}

	newNode := &Node{value, nil}

	if index == 0 {
		newNode.next = list.head
		list.head = newNode
	} else {
		current := list.head

		for i := 0; i < index-1; i++ {
			current = current.next
		}

		newNode.next = current.next
		current.next = newNode
	}

	list.size++
}

// Remove o nó contendo o valor dado da lista encadeada
func (list *LinkedList) Remove(value int) {
	if list.head == nil {
		panic("List is empty")
	}
	if list.head.value == value {
		list.head = list.head.next
		list.size--
		return
	}
	current := list.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next == nil {
		panic("Value not found")
	}
	current.next = current.next.next
	list.size--
}

// RemoveOnIndex remove o nó na posição index da lista encadeada
func (list *LinkedList) RemoveOnIndex(index int) {
	if index < 0 || index >= list.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		list.head = list.head.next
	} else {
		current := list.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	list.size--
}

// Get retorna o valor do nó na posição index da lista encadeada
func (list *LinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		panic("Index out of bounds")
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

// Set altera o valor do nó na posição index da lista encadeada para o valor dado
func (ll *LinkedList) Set(index int, value int) {
	if index < 0 || index >= ll.size {
		panic("Index out of bounds")
	}
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.value = value
}

// Size retorna o tamanho da lista encadeada
func (list *LinkedList) Size() int {
	return list.size
}

// Imprime a lista
func printLinkedList(list *LinkedList) {
	current := list.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	// Criando uma nova lista encadeada
	myLinkedList := LinkedList{}

	// Adicionando valores na lista
	myLinkedList.Add(3)
	myLinkedList.Add(7)
	myLinkedList.Add(1)
	myLinkedList.Add(8)

	// Imprimindo a lista
	fmt.Println("Lista encadeada inicial:")
	printLinkedList(&myLinkedList)

	// Adicionando um valor na posição 2 da lista
	myLinkedList.AddOnIndex(2, 5)

	// Imprimindo a lista após adicionar um valor na posição 2
	fmt.Println("\nLista encadeada após adicionar 5 na posição 2:")
	printLinkedList(&myLinkedList)

	// Removendo o valor 7 da lista
	myLinkedList.Remove(7)

	// Imprimindo a lista após remover o valor 7
	fmt.Println("\nLista encadeada após remover o valor 7:")
	printLinkedList(&myLinkedList)

	// Alterando o valor na posição 1 da lista para 9
	myLinkedList.Set(1, 9)

	// Imprimindo a lista após alterar o valor na posição 1 para 9
	fmt.Println("\nLista encadeada após alterar o valor na posição 1 para 9:")
	printLinkedList(&myLinkedList)

	// Imprimindo o valor na posição 3 da lista
	fmt.Println("\nValor na posição 3 da lista:", myLinkedList.Get(3))

	// Removendo o valor na posição 0 da lista
	myLinkedList.RemoveOnIndex(0)

	// Imprimindo a lista após remover o valor na posição 0
	fmt.Println("\nLista encadeada após remover o valor na posição 0:")
	printLinkedList(&myLinkedList)

	// Imprimindo o tamanho da lista
	fmt.Println("\nTamanho da lista encadeada:", myLinkedList.Size())
}
