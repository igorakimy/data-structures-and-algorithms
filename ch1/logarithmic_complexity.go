package main

import (
	"fmt"
)

type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

func printTree(tree *Tree) {
	if tree != nil {
		fmt.Println(" Value", tree.Value)
		fmt.Printf("Tree Node Left")
		printTree(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		printTree(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

func main() {
	var tree = &Tree{nil, 1, nil}
	printTree(tree)
	tree.insert(3)
	printTree(tree)
	tree.insert(5)
	printTree(tree)
	tree.LeftNode.insert(7)
	printTree(tree)
}

//  Value 1
// Tree Node LeftNil
// Tree Node RightNil
// Value 1
// Tree Node Left Value 3
// Tree Node LeftNil
// Tree Node RightNil
// Tree Node RightNil
// Value 1
// Tree Node Left Value 3
// Tree Node LeftNil
// Tree Node RightNil
// Tree Node Right Value 5
// Tree Node LeftNil
// Tree Node RightNil
// Value 1
// Tree Node Left Value 3
// Tree Node Left Value 7
// Tree Node LeftNil
// Tree Node RightNil
// Tree Node RightNil
// Tree Node Right Value 5
// Tree Node LeftNil
// Tree Node RightNil
