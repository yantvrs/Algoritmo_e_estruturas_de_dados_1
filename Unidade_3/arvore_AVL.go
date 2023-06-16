package main

import "fmt"

/*
  Árvores AVL
    AVL é uma árvore de busca binária balanceada com relação à altura de suas subárvores
      Uma árvore AVL verifica a altura das subárvores da esquerda e da direita, garantindo que essa diferença não seja maior que +/- 1
    Esta diferença é seu Fator de balanceamento
      bf = hr - hl
 
  Como balancear uma árvore? Algoritmo AVL

  . Inserção na subárvore direita do filho à direita
    > Solução: Rotação simples à esquerda
  . Inserção na subárvore esquerda do filho à esquerda
    > Solução: Rotação simples à direita
  . Inserção na subárvore esquerda do filho à direita
    > Solução: Rotação dupla direita-esquerda
  . Inserção na subárvore direita do filho à esquerda
    > Solução: Rotação dupla esquerda-direita
  
*/

type BstNode struct {
  left *BstNode
  value int
  height int
  bf int
  right *BstNode
}

func newNode(value int) *BstNode {
	return &BstNode{
		value: value,
	}
}

func (bstNode *BstNode) RotationRight() BstNode{
  left := bstNode.left
  bstNode.left = left.right
  left.right = bstNode
  bstNode.UpdateProperties()func newNode(value int) *BstNode {
	return &BstNode{
		value: value,
	}
}
  left.UpdateProperties()func (bstNode *BstNode) LeftLeft() *BstNode {
    return bstNode.RotationRight()
}
  return left
}

func (bstNode *BstNode) RotationLeft() BstNode{
  right := bstNode.right
  bstNode.right = right.left
  right.left = bstNode
  bstNode.UpdateProperties()
  right.UpdateProperties()
  return right
}

func (bstNode *BstNode) UpdateProperties() {
    // atualizar altura
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
            bstNode.height = 1 + heightRight
        } else {
            bstNode.height = 1 + heightLeft
        }
    }
    
    // atualizar bf
    bstNode.bf = heightRight - heightLeft
}

func (bstNode *BstNode) Add( value int ) {
  if value <= bstNode.value {
      if bstNode.left == nil {
          bstNode.left = NewNode(value:value)
      } else {
          bstNode.left.Add(value)
      }
  } else {
      if bstNode.right == nil {
          bstNode.right = NewNode(value:value)
      } else {
          bstNode.right.Add(value)
      }
  }
	
  bstNode.UpdateProperties()
  bstNode.Rebalance()
}

func (bstNode *BstNode) Rebalance() *BstNode {
    // Left
    if bstNode.bf <= -2 {
        // Left-left

func (bstNode *BstNode) RebalanceLeftLeft() *BstNode {
    return bstNode.RotationRight()
}

func (bstNode *BstNode) RebalanceLeftRight() *BstNode {
    bstNode.left = bstNode.left.RotationLeft()
    return bstNode.rebalanceLeftLeft()
}

func (bstNode *BstNode) RebalanceRightRight() *BstNode {
    return bstNode.RotationLeft()
}

func (bstNode *BstNode) rebalanceRightLeft() *BstNode {
    bstNode.right = bstNode.right.RotationRight()
    return bstNode.rebalanceRightRight()
}
