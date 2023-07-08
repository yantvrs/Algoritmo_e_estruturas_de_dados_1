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

	Retornar um ponteiro para um *BstNode na função de adição em árvores AVL é fundamental para refletir as alterações na estrutura da árvore original. Esse retorno permite atualizar corretamente a árvore após a adição de um novo nó, realizar a recursão adequada para encontrar a posição correta do nó e aplicar as ações de rebalanceamento necessárias para manter a propriedade de balanceamento da árvore AVL. Além disso, o retorno do ponteiro atualizado garante a consistência da árvore e a integridade das referências entre os nós.
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

// func max and min

func (bstNode *BstNode) Min() int {
	if bstNode.left != nil {
		return bstNode.left.Min()
	} else {
		return bstNode.value
	}
}

func (bstNode *BstNode) Max() int {
	if bstNode.right == nil {
		return bstNode.value
	} else {
		return bstNode.right.Max()
	}
}

func (bstNode *BstNode) Remove(value int) *BstNode {
	if value < bstNode.value {
		bstNode.left = bstNode.left.Remove(value)
	} else if value > bstNode.value {
		bstNode.right = bstNode.right.Remove(value)
	} else {
		if bstNode.right == nil && bstNode.left == nil {
			return nil 
		} else if bstNode.right != nil && bstNode.left == nil {
			return bstNode.right
		} else if bstNode.right == nil && bstNode.left != nil {
			return bstNode.left
		} else {
			maxValue := bstNode.left.Max()
			bstNode.value = maxValue
			bstNode.left = bstNode.left.Remove(maxValue)
			return bstNode
		}
	}
	bstNode.UpdateProperties()
	return bstNode.Rebalance()
}
