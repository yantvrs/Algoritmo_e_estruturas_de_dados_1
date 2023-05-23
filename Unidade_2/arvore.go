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

func (bstNode *BstNode) Min() int {
    if bstNode.left == nil {
        return bstNode.value
    } else {
        return bstNode.left.Min()
    }
}

func (bstNode *BstNode) Max() int {
    if bstNode.right == nil {
        return bstNode.value
    } else {
        return bstNode.right.Max()
    }
}

func (bstNode *BstNode) Println() {
   
    
}

// pre e pos

func (bstNode *BstNode) Height() int {
    htLeft = 0
    htRight = 0
    
    if bstNode.left != nil {
        htLeft = bstNode.left.Height()
    }
    if bstNode.right != nil {
        htRight = bstNode.right.Height()
    }
    
    if htLeft >= htRight {
        return 1 + htLeft
    } else {
        return 1 + htRight
    }
}

func main() {
    fmt.Println("Hello world!")
}


