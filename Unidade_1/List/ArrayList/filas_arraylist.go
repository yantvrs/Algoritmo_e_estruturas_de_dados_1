package main

import "fmt"

type Queue struct {
	elements []int
}

func (queue *Queue) Enqueue(e int) {
	queue.elements = append(queue.elements, e)
}

func (queue *Queue) Dequeue() int {
	if len(queue.elements) == 0 {
		fmt.Println("A fila está vazia")
		return -1
	}
	first := queue.elements[0]
	queue.elements = queue.elements[1:]
	return first
}

func (queue *Queue) Front() int {
	if len(queue.elements) == 0 {
		fmt.Println("A fila está vazia")
		return -1
	}
	return queue.elements[0]
}

func (queue *Queue) Size() int {
	return len(queue.elements)
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.elements) == 0
}

func main() {
	queue := &Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Front())
	fmt.Println(queue.Size())
	fmt.Println(queue.IsEmpty())
}
