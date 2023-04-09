package main

import "fmt"

// Estrutura que representa o nó
type Node struct {
	value int
	next  *Node
}

// Inicialização de variáveis representativas da lista
type LinkedList struct {
	head *Node
	size int
}

// Função que imprime a lista
func (list *LinkedList) Print() {

	auxiliary := list.head

	for auxiliary != nil {
		fmt.Println(auxiliary.value)
		auxiliary = auxiliary.next
	}
}

// Adiciona elemento no final
func (list *LinkedList) AddEnd(value int) {
	newNode := &Node{value, nil}

	auxiliary := list.head

	if auxiliary == nil {
		list.head = newNode
	} else {

		// Funciona como o while
		for auxiliary.next != nil {
			auxiliary = auxiliary.next
		}

		auxiliary.next = newNode
	}
	list.size++
}

// Adiciona em uma posição específica
func (list *LinkedList) AddIndex(value int, index int) {
	newNode := &Node{value: value, next: nil}

	if index < 0 || index > list.size {
		fmt.Println("O index não pertence ao intervalo da lista")
		return
	}
	if index == 0 {
		newNode.next = list.head
		list.head = newNode
	} else {
		auxiliary := list.head
		for i := 0; i < index-1; i++ {
			auxiliary = auxiliary.next
		}
		newNode.next = auxiliary.next
		auxiliary.next = newNode
	}
	list.size++
}

// Remove elemento em uma posição específica
func (list *LinkedList) RemoveIndex(index int) {
	if index < 0 || index >= list.size {
		fmt.Println("O index não pertence ao intervalo da lista")
		return
	}
	if index == 0 {
		list.head = list.head.next
	} else {
		auxiliary := list.head
		for i := 0; i < index-1; i++ {
			auxiliary = auxiliary.next
		}
		auxiliary.next = auxiliary.next.next
	}
	list.size--
}

// Obter elemento em uma dada posição
func (list *LinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		fmt.Println("O index não pertence ao intervalo da lista")
		return -1
	}
	auxiliary := list.head
	for i := 0; i < index; i++ {
		auxiliary = auxiliary.next
	}
	return auxiliary.value
}

// Trocar um elemento por outro
func (list *LinkedList) Set(newItem int, index int) {
	if index < 0 || index >= list.size {
		fmt.Println("O index não pertence ao intervalo da lista")
		return
	}
	auxiliary := list.head
	for i := 0; i < index; i++ {
		auxiliary = auxiliary.next
	}
	auxiliary.value = newItem
}

// Tamanho da lista
func (list *LinkedList) Size() int {
	return list.size
}

func main() {
	list := LinkedList{}

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
