package Sort

import "fmt"

func SelectionSort(array []int){
    size := len(v)
    for sweep := 0; sweep < size - 1; sweep++ {
        iSmaller := sweep
        for i := sweep; i < size - 1; sweep ++ {
            if array[i] < array[iSmaller] {
                iSmaller = i
            }
         array[sweep],array[iSmaller] = array[iSmaller],array[sweep]
        }
    }
}
