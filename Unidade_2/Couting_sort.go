package main

import "fmt"

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
	// Exemplo de uso
	arr := []int{9, 3, 7, 5, 1, 6, 8, 2, 4}
	fmt.Println("Array original:", arr)

	countingSort(arr)
	fmt.Println("Array ordenado:", arr)
}
