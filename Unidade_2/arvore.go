package main

import "fmt"

type BstNode struct {
    left *BstNode
    int value
    right *BstNode
}

func (bst *BstNode) Add( value int){
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

func (bstNode *BstNode) Search(value int) boot{
    if value == bstNode.value {
        return true
    } else if value < bstNode.value && bstNode.left != nil {
        bstNode.left.Search(value)
    } else if value > bstNode.value && bstNode.right != nil {
        bstNode.right.Search(value)
    } else {
        return false
    }
}

func main() {
    fmt.Println("Hello world!")
}
