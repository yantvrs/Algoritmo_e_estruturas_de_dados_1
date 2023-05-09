package main

import "fmt"

func MergeSort( v1 []int, v2 []int) []int{
	sizeV1 := len(v1)
	sizeV2 := len(v2)
	sizeV3 := sizeV1 + sizeV2

	v3 := make([]int, sizeV3)

	i:=0
	j:=0
	k:=0

	for i < sizeV1 && j < sizeV2 {
		if v2[j] < v1[i]{
			v3[k] = v2[j]
			k++
			j++
		} else {
			v3[k] = v1[i]
			k++
			i++
		}
	}
	if j > i {
		v3[sizeV3-1] = v1[sizeV1-1]
	} else {
		v3[sizeV3-1] = v2[sizeV2-1]
	}
	return v3
}


func main() {
    v1 := []int {5, 4, 2, 0, 1}
	v2 := []int{3, 10, 6, 8, 7}
    fmt.Println(v1,v2)
    
    fmt.Println(MergeSort(v1,v2))
}