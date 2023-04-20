package swap

func swap(array []int, i int, j int) {
	auxiliary := array[i]
	array[i] = array[j]
	array[j] = auxiliary
}
