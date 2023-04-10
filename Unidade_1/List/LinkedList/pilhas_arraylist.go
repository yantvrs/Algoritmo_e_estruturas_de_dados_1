package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (stack *Stack) Push(e int) {
	newNode := &Node{e, stack.top}
	stack.top = newNode
	stack.size++
}

func (stack *Stack) Pop() int {
	if stack.size == 0 {
		fmt.Println("A pilha está vazia")
		return -1
	}
	top := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return top
}

func (stack *Stack) Top() int {
	if stack.size == 0 {
		fmt.Println("A pilha está vazia")
		return -1
	}
	return stack.top.value
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Size())
	fmt.Println(stack.IsEmpty())
}
