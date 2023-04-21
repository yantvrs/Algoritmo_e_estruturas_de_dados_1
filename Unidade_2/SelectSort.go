package Sort

func SelectSort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		for j := i + 1; j < len(v); j++ {
			if v[i] > v[j] {
				swap(v, i, j)
			}
		}
	}
}
