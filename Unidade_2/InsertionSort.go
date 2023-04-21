package main

import "fmt"

func InsertionSort(v[] int){
    
    for i := 1; i < len(v); i++ {
        for j:= i; j > 0; j--{
            if v[j] < v[j-1] {
                v[j], v[j-1] = v[j-1], v[j]
            } else {
                break
            }
        }
    }
}


func main() {
    v := []int {5, 4, 2, 0, 1, 10, 6, 8, 7}
    fmt.Println(v)
    
    InsertionSort(v)
    fmt.Println(v)
}
