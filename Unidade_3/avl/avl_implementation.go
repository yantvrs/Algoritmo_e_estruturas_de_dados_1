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
	UpdateProperties()

	Retornar um ponteiro para um nó da árvore AVL nas funções de rebalanceamento é fundamental para garantir que as rotações realizadas durante o processo de balanceamento sejam refletidas na estrutura da árvore original. Ao retornar o ponteiro atualizado, as alterações nos relacionamentos entre os nós e a posição do nó desbalanceado são corretamente propagadas, mantendo a consistência da árvore. Isso é essencial para realizar operações como inserção, remoção e rebalanceamento de forma adequada e preservar as propriedades de uma árvore AVL.
*/

func (bstNode *BstNode) Rebalance() *BstNode {
	if bstNode.bf <= -2 {
		if bstNode.left.bf == -1 {
			return bstNode.RebalanceToLeftLeft()
		} else if bstNode.left.bf == 0 {
			return bstNode.RebalanceToLeftLeft()
		} else {
			return bstNode.RebalanceToLeftRight()
		}
	} else if bstNode.bf >= 2 {
		if bstNode.right.bf == 1 {
			return bstNode.RebalanceToRightRight()
		} else if bstNode.right.bf == 0 {
			return bstNode.RebalanceToRightRight()				
		} else {
			return bstNode.RebalanceToRightLeft()
		}
	}
	return bstNode
}

//Rotações 

func (bstNode *BstNode) RotRight() *BstNode{
	newRoot := bstNode.left
	bstNode.left = newRoot.right
	newRoot.right = bstNode

	// A ordem dos elementos a seguir é muito importante
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

// Rebalanceado árvore

func (bstNode *BstNode) RebalanceToLeftLeft() *BstNode {
	return bstNode.RotRight()
}

func (bstNode *BstNode) RebalanceToLeftRight() *BstNode {
	bstNode.left = bstNode.RotRight() // Redução do caso 2 para o caso 1
	return bstNode.Rotleft()
}

func (bstNode *BstNode) RebalanceToRightRight() *BstNode {
	return bstNode.Rotleft()
}

func (bstNode *BstNode) RebalanceToRightLeft() *BstNode {
	bstNode.right = bstNode.Rotleft() // Redução do caso 4 para o caso 3
	return bstNode.RotRight()
}

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
