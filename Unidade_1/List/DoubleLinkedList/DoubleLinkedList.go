package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (list *LinkedList) Print() {
	auxiliary := list.head
	for auxiliary != nil {
		fmt.Print(auxiliary.value, ",")
		auxiliary = auxiliary.next
	}
	println()
}

func (list *LinkedList) AddEnd(value int) {
	newNode := &Node{value, nil, nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.size++
}

func (list *LinkedList) AddIndex(value int, index int) {
	newNode := &Node{value: value, next: nil, prev: nil}

	if index < 0 || index > list.size {
		fmt.Println("O index não pertence ao intervalo da lista")
		return
	}
	if index == 0 {
		newNode.next = list.head
		if list.head != nil {
			list.head.prev = newNode
		}
		list.head = newNode
		if list.tail == nil {
			list.tail = newNode
		}
	} else if index == list.size {
		list.AddEnd(value)
		return
	} else {
		auxiliary := list.head
		for i := 0; i < index-1; i++ {
			auxiliary = auxiliary.next
		}
		newNode.next = auxiliary.next
		newNode.prev = auxiliary
		auxiliary.next.prev = newNode
		auxiliary.next = newNode
	}
	list.size++
}

func (list *LinkedList) RemoveIndex(index int) {
	if index < 0 || index >= list.size {
		fmt.Println("O index não pertence ao intervalo da lista")
		return
	}
	if index == 0 {
		list.head = list.head.next
		if list.head != nil {
			list.head.prev = nil
			if list.head.next == nil {
				list.tail = list.head
			}
		} else {
			list.tail = nil
			list.head = nil
			return
		}
	} else if index == list.size-1 {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
	} else {
		auxiliary := list.head
		for i := 0; i < index-1; i++ {
			auxiliary = auxiliary.next
		}
		auxiliary.next = auxiliary.next.next
		auxiliary.next.prev = auxiliary
	}
	list.size--
}

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

func (list *LinkedList) Size() int {
	return list.size
}

func main() {
	list := &LinkedList{}

	list.AddEnd(1)
	list.AddEnd(2)
	list.AddEnd(3)
	list.Print()
	fmt.Println()

	list.AddIndex(4, 1)
	list.Print()
	fmt.Println()

	list.RemoveIndex(2)
	list.Print()
	fmt.Println()

	fmt.Println(list.Get(1))

	list.Set(5, 1)
	list.Print()
}
