package main

// Node representa um nó da lista duplamente encadeada
type Node struct {
	value int
	prev  *Node
	next  *Node
}

// DoublyLinkedList representa uma lista duplamente encadeada
type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// Add adiciona um novo nó contendo o valor dado no final da lista duplamente encadeada
func (dll *DoublyLinkedList) Add(value int) {
	newNode := &Node{value, nil, nil}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
	dll.size++
}

// AddOnIndex adiciona um novo nó contendo o valor dado na posição index da lista duplamente encadeada
func (dll *DoublyLinkedList) AddOnIndex(index int, value int) {
	if index < 0 || index > dll.size {
		panic("Index out of bounds")
	}
	newNode := &Node{value, nil, nil}
	if index == 0 {
		if dll.head == nil {
			dll.head = newNode
			dll.tail = newNode
		} else {
			newNode.next = dll.head
			dll.head.prev = newNode
			dll.head = newNode
		}
	} else if index == dll.size {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	} else {
		current := dll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		newNode.prev = current
		current.next.prev = newNode
		current.next = newNode
	}
	dll.size++
}

// Remove remove o nó contendo o valor dado da lista duplamente encadeada
func (dll *DoublyLinkedList) Remove(value int) {
	if dll.head == nil {
		panic("List is empty")
	}
	if dll.head.value == value {
		if dll.size == 1 {
			dll.head = nil
			dll.tail = nil
		} else {
			dll.head = dll.head.next
			dll.head.prev = nil
		}
RemoveOnIndex(index int) {
	if index < 0 || index >= dll.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		if dll.size == 1 {
			dll.head = nil
			dll.tail = nil
		} else {
			dll.head = dll.head.next
			dll.head.prev = nil
		}
	} else if index == dll.size-1 {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	} else {
		current := dll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
		current.next.prev = current
	}
	dll.size--
}

// Get retorna o valor do nó na posição index da lista duplamente encadeada
func (dll *DoublyLinkedList) Get(index int) int {
	if index < 0 || index >= dll.size {
		panic("Index out of bounds")
	}
	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}
RemoveOnIndex(index int) {
	if index < 0 || index >= dll.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		if dll.size == 1 {
			dll.head = nil
			dll.tail = nil
		} else {
			dll.head = dll.head.next
			dll.head.prev = nil
		}
	} else if index == dll.size-1 {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	} else {
		current := dll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
		current.next.prev = current
	}
	dll.size--
}

// Get retorna o valor do nó na posição index da lista duplamente encadeada
func (dll *DoublyLinkedList) Get(index int) int {
	if index < 0 || index >= dll.size {
		panic("Index out of bounds")
	}
	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

func main() {
	// Cria uma nova lista duplamente encadeada
	dll := &DoublyLinkedList{}

	// Adiciona alguns valores na lista
	dll.Add(1)
	dll.Add(2)
	dll.Add(3)

	// Imprime o tamanho da lista
	fmt.Println(dll.size) // 3

	// Adiciona um valor na posição 1 da lista
	dll.AddOnIndex(1, 4)

	// Imprime o tamanho da lista
	fmt.Println(dll.size) // 4

	// Imprime o valor na posição 1 da lista
	fmt.Println(dll.Get(1)) // 4

	// Remove o valor 4 da lista
	dll.Remove(4)

	// Imprime o tamanho da lista
	fmt.Println(dll.size) // 3

	// Remove o valor na posição 1 da lista
	dll.RemoveOnIndex(1)

	// Imprime o tamanho da lista
	fmt.Println(dll.size) // 2

	// Imprime o valor na posição 0 da lista
	fmt.Println(dll.Get(0)) // 1

	// Imprime o valor na posição 1 da lista
	fmt.Println(dll.Get(1)) // 3
}