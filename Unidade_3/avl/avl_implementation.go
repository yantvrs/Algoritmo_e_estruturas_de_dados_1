package main

import "fmt"

type BstNode struct {
	left *BstNode
	value int 
	height int 
	bf int 
	right *BstNode
}