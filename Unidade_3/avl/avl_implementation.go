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


/*
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
	Rebalance() *BstNode
	UpdateProperties()

	O retorno do tipo *BstNode na função RotRight() é fundamental para permitir que a rotação à direita altere a estrutura da árvore corretamente, atualizando os ponteiros dos nós. Ele permite que a nova raiz da subárvore seja retornada e atribuída a variáveis ou usada em outras operações, possibilitando o uso e manipulação adequados da árvore rotacionada.
*/

func (bstNode *BstNode) RotRight() *BstNode{
	newRoot := bstNode.left
	bstNode.left = newRoot.right
	newRoot.right = bstNode
	bstNode.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

func (bstNode *BstNode) Rotleft() *BstNode{
	newRoot := bstNode.right
	bstNode.right = newRoot.left
	newRoot.left = bstNode
	bstNode.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}



func (bstNode *BstNode) Add(value int) *BstNode {
	if value < bstNode.value {
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
	
func (bstNode *BstNode) UpdateProperties() {
	heightRight := 0
	heightLeft := 0

	if bstNode.right == nil && bstNode.left == nil {
		bstNode.height = 0 
	} else {
		if bstNode.right != nil {
			heightRight = bstNode.right.height
		}
		if bstNode.left != nil {
			heightLeft = bstNode.left.height
		}
		if heightRight > heightLeft {
			bstNode.height = heightRight + 1
		} else {
			bstNode.height = heightLeft + 1
		}
	}
	bstNode.bf = heightRight - heightLeft
}
