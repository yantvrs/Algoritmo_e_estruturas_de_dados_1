package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
	If
A declaração if parece com os laços for; a expressão não precisa ser cercada de ( ) mas os chaves { } são obrigatórios.
*/
