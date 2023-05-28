package main

import "fmt"

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

func newNode(value int) *BstNode {
	return &BstNode{
		value: value,
	}
}

func (bstNode *BstNode) Add(value int) {
	if value <= bstNode.value {
		if bstNode.left == nil {
			bstNode.left = newNode(value)
		} else {
			bstNode.left.Add(value)
		}
	} else {
		if bstNode.right == nil {
			bstNode.right = newNode(value)
		} else {
			bstNode.right.Add(value)
		}
	}
}

func (bstNode *BstNode) Search(value int) bool {
	if value == bstNode.value {
		return true
	} else if value < bstNode.value && bstNode.left != nil {
		return bstNode.left.Search(value)
	} else if value > bstNode.value && bstNode.right != nil {
		return bstNode.right.Search(value)
	} else {
		return false
	}
}

func (bstNode *BstNode) Min() int {
	if bstNode.left == nil {
		return bstNode.value
	}
	return bstNode.left.Min()
}

func (bstNode *BstNode) Max() int {
	if bstNode.right == nil {
		return bstNode.value
	}
	return bstNode.right.Max()
}

// raiz, esquerda, direita
func (bstNode *BstNode) PrintPre() {
	if bstNode == nil {
		return
	}
	fmt.Printf("%d ", bstNode.value)
	bstNode.left.PrintPre()
	bstNode.right.PrintPre()
}

// esquerda, raiz, direita
func (bstNode *BstNode) PrintIn() {
	if bstNode == nil {
		return
	}
	bstNode.left.PrintIn()
	fmt.Printf("%d ", bstNode.value)
	bstNode.right.PrintIn()
}

// esquerda, direita, raiz
func (bstNode *BstNode) PrintPos() {
	if bstNode == nil {
		return
	}
	bstNode.left.PrintPos()
	bstNode.right.PrintPos()
	fmt.Printf("%d ", bstNode.value)
}

func (bstNode *BstNode) Height() int {
	htLeftBasis := 0
	htRightBasis := 0

	if bstNode.left != nil {
		htLeftBasis = 1 + bstNode.left.Height()
	}
	if bstNode.right != nil {
		htRightBasis = 1 + bstNode.right.Height()
	}

	if htLeftBasis >= htRightBasis {
		return htLeftBasis
	} else {
		return htRightBasis
	}
}

// Caso 1: Remoção no nó folha
// Caso 2: Nó com um filho
// Caso 3 : Remoção de nó com dois filhos

func (bstNode *BstNode) Remove(value int) *BstNode {
	if value < bstNode.value {
		bstNode.left = bstNode.left.Remove(value)
	} else if value > bstNode.value {
		bstNode.right = bstNode.right.Remove(value)
	} else {
		if bstNode.left == nil && bstNode.right == nil {
			return nil
		} else if bstNode.left != nil && bstNode.right == nil {
			return bstNode.left
		} else if bstNode.left == nil && bstNode.right != nil {
			return bstNode.right
		} else {
			bstNode.value = bstNode.left.Max()
			bstNode.left = bstNode.left.Remove(bstNode.left.Max())
			return bstNode
		}
	}
	return bstNode
}

func main() {
	fmt.Println("Binary Search Tree Example:")

	// Criação da árvore
	root := newNode(5)
	root.Add(3)
	root.Add(8)
	root.Add(2)
	root.Add(4)
	root.Add(7)
	root.Add(9)

	// Teste das funções
	fmt.Println("Preorder Traversal:")
	root.PrintPre()
	fmt.Println()

	fmt.Println("Inorder Traversal:")
	root.PrintIn()
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	root.PrintPos()
	fmt.Println()

	fmt.Println("Minimum value:", root.Min())
	fmt.Println("Maximum value:", root.Max())

	fmt.Println("Height of the tree:", root.Height())

	value := 4
	fmt.Println("Search for value", value, ":", root.Search(value))

	valueToRemove := 7
	fmt.Println("Removing value", valueToRemove)
	root = root.Remove(valueToRemove)

	fmt.Println("Inorder Traversal after removal:")
	root.PrintIn()
	fmt.Println()
}
