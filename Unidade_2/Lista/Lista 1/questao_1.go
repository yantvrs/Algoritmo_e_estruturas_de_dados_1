package main

import (
	"fmt"
	"math/rand"
)

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


func partition(v []int, low int, high int) int {
	// Selecionar um pivô aleatório
	randomIndex := rand.Intn(high-low+1) + low
	v[high], v[randomIndex] = v[randomIndex], v[high]

	pivot := v[high]
	i := low - 1

	for j := low; j < high; j++ {
		if v[j] <= pivot {
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	v[i+1], v[high] = v[high], v[i+1]
	return i + 1
}

func quicksort(v []int, low int, high int) {
	if low < high {
		// Encontrar o índice de partição
		pivot := partition(v, low, high)

		// Ordenar recursivamente as duas metades
		quicksort(v, low, pivot-1)
		quicksort(v, pivot+1, high)
	}

}

// f)CoutingSort
func countingSort(v []int) {
	sizeV := len(v)

	menor := v[0]
	maior := v[0]
	
	for i := 1; i < sizeV; i++ {
		if v[i] < menor {
			menor = v[i]
		}
		if v[i] > maior {
			maior = v[i]
		}
	}

	novoTam := maior - menor + 1

	contagem := make([]int, novoTam)

	for i := 0; i < sizeV; i++ {
		contagem[v[i]-menor]++
	}

	for i := 1; i < novoTam; i++ {
		contagem[i] += contagem[i-1]
	}

	ordenado := make([]int, sizeV)

	for i := 0; i < sizeV; i++ {
		ordenado[contagem[v[i]-menor]-1] = v[i]
	}

	for i := 0; i < sizeV; i++ {
		v[i] = ordenado[i]
	}
}


func main() {
	fmt.Println("Algoritmos de ordenação")

	array := []int{9,8,6,5,3,2,0,4,1,7}
	fmt.Println("Vetor desordenado:", array)

	alternativa := "e"

	if alternativa == "a" {
		// a) SelectionSort
		fmt.Println("\nSelectionSort:")
		SelectionSort(array)
		fmt.Println("Ordenado:", array)
	}
	if alternativa == "b" {
		// b) BubbleSort
		fmt.Println("\nBubbleSort:")
		BubbleSort(array)
		fmt.Println("Ordenado:", array)
	}
	if alternativa == "c" {
		// c) InsertionSort
		fmt.Println("\nInsertionSort:")
		InsertionSort(array)
		fmt.Println("Ordenado:", array)
	}
	if alternativa == "d" {
		// d) MergeSort
		fmt.Println("\nMergeSort:")
		MergeSort(array, 0, len(array)-1)
		fmt.Println("Ordenado:", array)
	}
	if alternativa == "e" {
		// e) QuickSort
		fmt.Println("\nQuickSort:")
		quicksort(array, 0, len(array)-1)
		fmt.Println("Ordenado:", array)
	}
	if alternativa == "f" {
		// f) CountingSort
		fmt.Println("\nCountingSort:")
		countingSort(array)
		fmt.Println("Ordenado:", array)
	}
}