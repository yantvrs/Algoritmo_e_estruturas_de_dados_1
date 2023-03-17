package main

import "fmt"

func main() {
	s := []int{2, 3, 8, 9, 6, 5, 4}

	s = s[1:5]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = s[:]
	fmt.Println(s)
}

/*
Padrões de uma Slice
Ao "fatiar", você pode omitir as posições de limite altas ou baixas para usar seus padrões em seu lugar.

O padrão é igual a zero para o limite baixo e o comprimento da slice para o limite alto.

Para a matriz

var a [10]int
essas expressões das slices são equivalentes:

a[0:10]
a[:10]
a[0:]
a[:]
*/
