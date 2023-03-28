package main

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

// Remove remove o nó contendo o valor dado da lista encadeada
func (ll *LinkedList) Remove(value int) {
	if ll.head == nil {
		panic("List is empty")
	}
	if ll.head.value == value {
		ll.head = ll.head.next
		ll.size--
		return
	}
	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next == nil {
		panic("Value not found")
	}
	current.next = current.next.next
	ll.size--
}

// RemoveOnIndex remove o nó na posição index da lista encadeada
func (ll *LinkedList) RemoveOnIndex(index int) {
	if index < 0 || index >= ll.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		ll.head = ll.head.next
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	ll.size--
}

// Get retorna o valor do nó na posição index da lista encadeada
func (ll *LinkedList) Get(index int) int {
	if index < 0 || index >= ll.size {
		panic("Index out of bounds")
	}
	current := ll.head
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
