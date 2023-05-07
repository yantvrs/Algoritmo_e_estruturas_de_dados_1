package main

import "fmt"

func BubbleSort(v []int) {
	size := len(v)
	for j := 0; j < size-1; j++ {
		change := false
		for i := 0; i < size-1-j; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				change = true
			}
		}
		if change == false {
			break
		}
	}
}

func main() {
	array := []int{9, 2, 5, 6, 7, 8, 4, 1, 3}

	fmt.Println("Unordered vector", array)

	BubbleSort(array)
	fmt.Println("Ordered", array)
	
}
