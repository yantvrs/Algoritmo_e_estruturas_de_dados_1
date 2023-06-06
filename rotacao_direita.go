package main

import "fmt"

/*
 
  Fator de balanceamento é calculado através da diferença entre a altura da direita menos o da esquerda.
  bf = hr - hl
  
  Desbalanceado: 
  bf >= 2
  bf <= -2 
  
  Caso 1: Desbalanceado para a esquerda -> bf= -2 -> Rotaciona p/ direita
*/

type BstNode struct {
  left *BstNode
  value int
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


