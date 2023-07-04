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

func NewNode(value int) *BstNode {
	node := &BstNode{left: nil, value: value, height: 0, bf: 0, right: nil}
	return node
}

func (bstNode *BstNode) 

func (bstNode *BstNode) Add(value int) *BstNode {
	if value <= bstNode.value {
		if bstNode.left == nil {
			bstNode.left = NewNode(value)
		} else {
			bstNode.left = bstNode.left.Add(value)
		}
	} else {
		if bstNode.right == nil {
			bstNode.right = NewNode(value)
		} else {
			bstNode.right = bstNode.right.Add(value)
		}
	}
	bstNode.UpdateProperties()

	return bstNode.Rebalance()
}
