package main

import (
	"fmt"
	"math/rand"
)

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

func main() {
	// Exemplo de uso
	arr := []int{9, 3, 7, 5, 1, 6, 8, 2, 4}
	fmt.Println("Array original:", arr)

	quicksort(arr, 0, len(arr)-1)
	fmt.Println("Array ordenado:", arr)
}
