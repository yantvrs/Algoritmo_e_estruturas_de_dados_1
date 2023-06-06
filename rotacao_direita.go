package main

import "fmt"

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


