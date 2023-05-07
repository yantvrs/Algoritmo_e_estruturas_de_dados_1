package main

import "fmt"

func SelectionSort(arr []int) {
    size := len(arr)
    for i := 0; i < size-1; i++ {
        iSmaller := i
        for j := i + 1; j < size; j++ {
            if arr[j] < arr[iSmaller] {
                iSmaller = j
            }
        }
        if iSmaller != i {
            arr[i], arr[iSmaller] = arr[iSmaller], arr[i]
        }
    }
}

func main() {
	array := []int{9, 2, 5, 6, 7, 8, 4, 1, 3}

	fmt.Println("Unordered vector", array)

	SelectionSort(array)
	fmt.Println("Ordered", array)
	
}
