package main

import "fmt"

/*
  Árvores AVL
    AVL é uma árvore de busca binária balanceada com relação à altura de suas subárvores
      Uma árvore AVL verifica a altura das subárvores da esquerda e da direita, garantindo que essa diferença não seja maior que +/- 1
    Esta diferença é seu Fator de balanceamento
    bf = hr - hl
 
  
  Desbalanceado: 
  bf >= 1
  bf <= -1
  
Casos:
  Raiz        Raiz.left       Ação
              (levemente)
> Esquerda -> Esquerda    ->  raiz.RotDir()
  bf = -2     bf = -1
> Esquerda -> Balanceada  ->  raiz.RotDir()
  bf = -2     bf = 1
> Esquerda -> Direita     ->  caso 1: raiz.left.RotEsq()
                              caso 2: raiz.RotDir()

  Raiz        Raiz.right
              (levemente)
> Direita  -> Direita    ->  raiz.RotEsq()
  bf = 2      bf = 1
> Direita  -> Balanceada ->  raiz.RotEsq()
  bf = 2      bf = 0
> Direita  -> Esquerda   ->  caso 1: raiz.Right.RotDir()
                             caso 2: raiz.RotEsq()
  
  
  >> raiz.bf == -2 && raiz.left.bf <= 0 
     Levemente desbalanceada para a esquerda ou completa
     ação: raiz.rot.Dir()
  
*/

type BstNode struct {
  left *BstNode
  value int
  height int
  bf int
  right *BstNode
}

func (bstNode *BstNode) UpdateProperties() {
    // atualizar altura
    heightRight := 0 
    heightLeft := 0
    if bstNode.right != nil && bstNode.left != nil {
        bstNode.height = 0 
    }
    else {
        if bstNode.right != nil {
            heightRight = bstNode.right.height
        }
        if bstNode.left != nil {
            heightLeft = bstNode.left.height
        }
        if heightRight > heightLevel {
            bstNode
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


