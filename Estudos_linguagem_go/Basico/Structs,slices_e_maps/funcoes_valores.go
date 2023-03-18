package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

/*
	Funções valores
Funções são valores também. Elas podem ser passadas assim como outros valores.

Funções valores podem ser usadas como argumentos de funções e retornar valores.
*/
