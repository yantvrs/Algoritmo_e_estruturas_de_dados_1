package main

type linkedList struct {
	head *firstNode
	size int
}

type firstNode struct {
	value    int
	nextNode *firstNode
}

func (list *linkedList) Add(value int) {
	newNode := firstNode{value, nil}

	aux := list.head
	prev := aux

	for aux != nil {
		prev = aux
		aux = aux.nextNode
	}

	if prev == nil {
		list.head = &newNode
	} else {
		prev.nextNode = &newNode
	}

	list.size++
}
