package main

// Estrutura que representa a linkedlist
type linkedList struct {
	head *node
	size int
}

// Estrutura que representa cada nó
type node struct {
	value int
	next  *node
}

// Adiciona elemento no final da lista
func (list *linkedList) Add(value int) {
	newNode := node{value, nil}

	assistant := linkedList.head
	prev := assistant

	for assistant != nil {
		prev := assistant
		assistant := assistant.next
	}

	if prev == nil {
		linkedList.head = &newNode
	} else {
		prev.next := &newNode
	}

	linkedList.size++
}
