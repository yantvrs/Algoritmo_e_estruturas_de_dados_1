package main

import "fmt"

/*1) Programe na linguagem C/Go cada um dos seguintes algoritmos de ordenação:*/

// a)SelectionSort (in-place)
func SelectionSort(v []int) {
	size := len(v)

	for i := 0 ; i < size - 1; i++ {
		iSmaller := i

		for j := i + 1; j < size ; j++ {
			if v[j] < v[iSmaller] {
				iSmaller = j
			}
		}

		if iSmaller != i {
			v[i],v[iSmaller]=v[iSmaller],v[i]
		}
	}
}

// b)BubbleSort (in-place)
func BubbleSort(v []int) {
	size := len(v)

	for j := 0; j < size - 1; j++ {
		change := false
		for i := 0 ; i < size - j - 1; i++ {
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

// c)InsertionSort (in-place)
func InsertionSort(v []int) {
	size := len(v)
	for i := 1 ; i < size ; i++ {
		for j := i; j > 0; j--{
			if v[j] < v[j-1] {
				v[j], v[j-1] = v[j-1], v[j]
			} else {
				break
			}
		}
	} 
}

// d)MergeSort
func Merge(v []int, start int, middle int, end int) {

	sizeLeft := middle - start + 1
	sizeRight := end - middle

	left := make([]int,sizeLeft)
	right := make([]int,sizeRight)

	for i := 0; i < sizeLeft; i++ {
		left[i] = v[start + i]
	}
	for i := 0; i < sizeRight; i++ {
		right[i] = v[middle + 1 + i]
	}
	
	i := 0
	j := 0
	k := start

	for i < sizeLeft && j < sizeRight {
		if left[i] < right[j]{
			v[k] = left[i]
			i++
		} else {
			v[k] = right[j]
			j++
		}
		k++
	}

	for i < sizeLeft {
        v[k] = left[i]
        i++
        k++
    }
    for j < sizeRight {
        v[k] = right[j]
        j++
        k++
    }

}

func MergeSort(v []int, start int, end int) {
    if start < end {
        middle := (start + end) / 2
        MergeSort(v, start, middle)
        MergeSort(v, middle+1, end)
        Merge(v, start, middle, end)
    }
}

// e)QuickSort(com randomização de pivô)
func QuickSort() {

}

// f)CoutingSort
func CoutingSort() {

}

func main() {
	fmt.Println("Algoritmos de ordenação\nVetor desordenado:")
	
	array := []int{9, 2, 5, 6, 7, 8, 4, 1, 3}
	fmt.Println("\nVetor desordenado:", array)

	
	alternativa := "d"


	if alternativa == "a" {
		// a)
		fmt.Println("\nSelectionSort:")
		SelectionSort(array)
		fmt.Println("Ordenado: ", array)
	}
	if alternativa == "b" {
		// b) 
		fmt.Println("\nBubbleSort:")
		BubbleSort(array)
		fmt.Println("Ordenado: ", array)
	}
	if alternativa == "c" {
		// c) 
		fmt.Println("\nInsertionSort:")
		InsertionSort(array)
		fmt.Println("Ordenado: ", array)
	}
	if alternativa == "d" {
		// d) 
		fmt.Println("\nMergeSort:")
		MergeSort(array, 0, len(array)-1)
		fmt.Println("Ordenado: ", array)
	}

}