package main

import "fmt"

type Bst struct{
    root *nodeBst
    size int
}

struct nodeBst {
    left *nodeBst
    int value
    right *nodeBst
}

func newNode(value int) *nodeBst{
    node := nodeBst{}
    node.value = value
    node.left = nil
    node.right = nil
    return node
}

func (bst *Bst) Add( value int) {
    if bst.root == nil{
        bst.root = newNode(value)
    }
    auxiliary := bst.root
    prev := auxiliary
    
    for auxiliary != nil {
        prev = auxiliary
        if value <= auxiliary.value {
            auxiliary = auxiliary.left
        } else {
            auxiliary = auxiliary.right
        }
    }
}
