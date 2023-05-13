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
func InsertionSort() {

}

// d)MergeSort
func MergeSort() {

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

	// a)
	fmt.Println("\nSelectionSort:")
	SelectionSort(array)
	fmt.Println("Ordenado: ", array)

	// b) 
	fmt.Println("\nBubbleSort:")
	BubbleSort(array)
}