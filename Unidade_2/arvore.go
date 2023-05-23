package main

import "fmt"

/*
Add(val int)
Search(val int) bool
Min() int 
Max() int 
PrintPre()
PrintIn()
PrintPos()
Height() int 
Remove(val int) *BstNode
*/

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
    htLeftBasis = 0
    htRightBasis = 0
    
    if bstNode.left != nil {
        htLeftBasis = bstNode.left.Height()
    }
    if bstNode.right != nil {
        htRightBasis = bstNode.right.Height()
    }
    
    if htLeftBasis >= htRightBasis {
        return htLeftBasis
    } else {
        return htRightBasis
    }
}

func main() {
    fmt.Println("Hello world!")
}


