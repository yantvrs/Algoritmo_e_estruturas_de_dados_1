package main

import "fmt"

func countingSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	// Encontrar o valor máximo no array
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Criar um array de contagem e inicializá-lo com zeros
	count := make([]int, max+1)

	// Contar a ocorrência de cada elemento
	for _, num := range arr {
		count[num]++
	}

	// Atualizar o array de contagem com as posições corretas
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// Criar um array temporário para armazenar os elementos ordenados
	temp := make([]int, len(arr))

	// Colocar os elementos na posição correta no array temporário
	for _, num := range arr {
		temp[count[num]-1] = num
		count[num]--
	}

	// Copiar os elementos ordenados de volta para o array original
	for i, num := range temp {
		arr[i] = num
	}
}

func main() {
	// Exemplo de uso
	arr := []int{9, 3, 7, 5, 1, 6, 8, 2, 4}
	fmt.Println("Array original:", arr)

	countingSort(arr)
	fmt.Println("Array ordenado:", arr)
}
