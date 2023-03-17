package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

/*

Slices literais
Uma slice literal é como um array literal, sem o comprimento.

Esta é uma matriz literal:

[3]bool{true, true, false}
E isso cria a mesma matriz como acima, em seguida, constrói uma slice que faz referência a ela:

[]bool{true, true, false}
*/
