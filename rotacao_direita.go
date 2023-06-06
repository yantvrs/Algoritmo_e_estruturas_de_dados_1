package main

import "fmt"

/*
 
  Fator de balanceamento é calculado através da diferença entre a altura da direita menos o da esquerda.
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

func (bstNode *BstNode) RotationRight() BstNode{
  left := bstNode.left
  bstNode.left = bstNode.left.right
  left.right = bstNode
  return left
}

func (bstNode *BstNode) RotationLeft() BstNode{
  right := bstNode.right
  bstNode.right = bstNode.right.left
  right.left = bstNode
  return right
}


