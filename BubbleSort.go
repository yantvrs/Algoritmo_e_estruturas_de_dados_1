package main

import "fmt"



func BubbleSort( v []int) {
    size := len(v)
    for j := 0; j < size - 1; j++ {
        var change = false
        for i := 0; i < size - 1 -j; i++ {
            if v[i] > v[i+1] {
                v[i], v[i+1] = v[i+1], v[i]
                change = true
            }
        }
        if change != true {
            return
        }
    }
}

func main() {
    v = []int{6,9,8,7,4}
    BubbleSort(v)
    fmt.Print(v)
}
