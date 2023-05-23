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
        htLeftBasis = 1 + bstNode.left.Height()
    }
    if bstNode.right != nil {
        htRightBasis = 1 + bstNode.right.Height()
    }
    
    if htLeftBasis >= htRightBasis {
        return htLeftBasis
    } else {
        return htRightBasis
    }
}

/*
Três casos de remoção:
> Caso 1(Nó folha) : Atualizar ponteiro que apontava para o nó a ser deletado para nil, ou seja, a remoção será feita no nó folha.
> Caso 2(Nó com um único filho) : Atualizar ponteiro que apontava para o nó a ser deletado para o único filho daquele nó.
> Caso 3(Nó com dois filhos) : Atualizar ponteiro que apontava para o nó a ser deletado
*/

func (bstNode *BstNode) Remove(value int) *BstNode {
    if value < bstNode.value {
        bstNode.left = bstNode.left.Remove(value)
    } else if value > bstNode.value {
        bstNode.right = bstNode.right.Remove(value)
    } else { 
        // Encontramos o nó a ser removido
        // Caso 1 
        
        if bstNode.left == nil && bstNode.right == nil { 
            return nil
        }
        
        // Caso 2
        
        else if bstNode.left != nil && bstNode.right == nil {
            return bstNode.left
        }
        else if bstNode.left == nil && bstNode.right != nil {
            return bstNode.right
        }
        
        // Caso 3
        else {
            bstNode.value = bstNode.left.Max()
            return bstNode.left.Remove(bstNode.left.Max())
        }
    }
}

func main() {
    fmt.Println("Hello world!")
}


