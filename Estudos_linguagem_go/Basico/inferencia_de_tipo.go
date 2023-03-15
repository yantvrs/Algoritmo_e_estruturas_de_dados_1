package main

import "fmt"

func main() {
	v := 42 //change me!
	fmt.Printf("v is of type %T", v)
}

/*
	Inferência de tipo
Ao declarar uma variável, sem especificar o seu tipo (usando var sem um tipo ou na sintaxe :=), o tipo da variável é inferida a partir do valor do lado direito.

Quando o lado direito da declaração é digitado, a nova variável é do mesmo tipo:

var i int
j := i // j é um int
Mas, quando o lado direito contém uma constante numérica não tipificada, a nova variável pode ser um int, float64, ou complex128 dependendo da precisão da constante:

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
Tente alterar o valor inicial de v no código de exemplo e observar como seu tipo é afetado.
*/
