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

func (bstNode *BstNode) RotationRight() BstNode{
  left := bstNode.left
  bstNode.left = left.right
  left.right = bstNode
  bstNode.UpdateProperties()
  left.UpdateProperties()
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
  
}
