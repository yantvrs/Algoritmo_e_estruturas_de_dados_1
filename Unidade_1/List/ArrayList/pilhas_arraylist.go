package main

import "fmt"

type Stack struct {
	elements []int
}

func (stack *Stack) Push(e int) {
	stack.elements = append(stack.elements, e)
}

func (stack *Stack) Pop() int {
	if len(stack.elements) == 0 {
		fmt.Println("A pilha está vazia")
		return -1
	}
	top := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return top
}

func (stack *Stack) Top() int {
	if len(stack.elements) == 0 {
		fmt.Println("A pilha está vazia")
		return -1
	}
	return stack.elements[len(stack.elements)-1]
}

func (stack *Stack) Size() int {
	return len(stack.elements)
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.elements) == 0
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
