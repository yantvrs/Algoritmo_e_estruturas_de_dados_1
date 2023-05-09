package main

import "fmt"

/*
[6,2,9,8,7]
size = 5
j = 0
change = false
i = 0 
array[i] = 6
array[i+1] = 

*/

func BubbleSort(array []int) {
	size := len(array)
	for j := 0; j < size-1; j++ {
		change := false
		for i := 0; i < size-1-j; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
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
