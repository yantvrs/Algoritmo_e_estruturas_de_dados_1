package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt(2))
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.

	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*

Interfaces
Um tipo de interface é definida por um conjunto de métodos.

Um valor de tipo de interface pode conter qualquer valor que implementa esses métodos.

Nota: Tem um erro no exemplo na linha 22. Vertex (do tipo valor) não satisfaz Abser porque o método Abs é definido apenas em *Vertex (do tipo ponteiro)
*/
