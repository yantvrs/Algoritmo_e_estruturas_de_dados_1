package main

type ITree interface {
	Add(value int) *BstNode
	Search(value int) bool
	Min() int
	Max() int
	PrintPre()
	PrintIn()
	PrintPos()
	Height() int
	Remove(value int) *BstNode
	IsBst() bool
	Size() int
	RotRight() *BstNode
	RotLeft() *BstNode
	Rebalance() *BstNode
	UpdateProperties()
}

type BstNode struct {
	left   *BstNode
	value  int
	height int
	bf     int
	right  *BstNode
}

func (bstNode *BstNode) RotRight() *BstNode {

}
