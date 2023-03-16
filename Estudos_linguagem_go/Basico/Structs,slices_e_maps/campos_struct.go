package main

import "fmt"

type Number struct {
	X int
	Y int
}

func main() {
	v := Number{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)
}
