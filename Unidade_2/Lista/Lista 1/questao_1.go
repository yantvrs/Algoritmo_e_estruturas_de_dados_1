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


func partition(arr []int, low, high int) int {
	// Selecionar um pivô aleatório
	randomIndex := rand.Intn(high-low+1) + low
	arr[high], arr[randomIndex] = arr[randomIndex], arr[high]

	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quicksort(arr []int, low, high int) {
	if low < high {
		// Encontrar o índice de partição
		pivot := partition(arr, low, high)

		// Ordenar recursivamente as duas metades
		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
	}

}

// f)CoutingSort
func countingSort(v []int) {
	sizeV := len(v)

	iMenor := 0
	iMaior := 0

	for i := 0; i < sizeV; i++ {
		if v[i] < v[iMenor] {
			iMenor = i
		}
		if v[i] > v[iMaior] {
			iMaior = i
		}
	}

	novoTam := v[iMaior] - v[iMenor] + 1

	contagem := make([]int, novoTam)

	for i := 0; i < sizeV; i++ {
		indiceContagem := v[i] - v[iMenor]
		contagem[indiceContagem]++
	}

	for i := 1; i < novoTam; i++ {
		contagem[i] += contagem[i-1]
	}

	ordenado := make([]int, sizeV)
	adicionado := make([]bool, sizeV)

	for i := 0; i < sizeV; i++ {
		indiceOrdenado := contagem[v[i]-v[iMenor]] - 1
		for adicionado[indiceOrdenado] {
			indiceOrdenado--
		}
		ordenado[indiceOrdenado] = v[i]
		adicionado[indiceOrdenado] = true
	}

	for i := 0; i < sizeV; i++ {
		v[i] = ordenado[i]
	}
}


func main() {
	fmt.Println("Algoritmos de ordenação")

	array := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println("Vetor desordenado:", array)

	alternativa := "c"

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