package Sort

func BubbleSort(v []int) {
	size := len(v)
	for j := 0; j < size-1; j++ {
		change := false
		for i := 0; i < size-1-j; i++ {
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
