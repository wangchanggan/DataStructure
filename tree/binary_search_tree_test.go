package tree

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	binarySearchTree := NewBinarySearchTree(4)
	binarySearchTree.Left = NewBinarySearchTree(2)
	binarySearchTree.Right = NewBinarySearchTree(6)
	binarySearchTree.Left.Left = NewBinarySearchTree(1)
	binarySearchTree.Left.Right = NewBinarySearchTree(3)
	binarySearchTree.Right.Left = NewBinarySearchTree(5)
	binarySearchTree.Right.Right = NewBinarySearchTree(7)
	fmt.Println(binarySearchTree.Find(2))
	binarySearchTree.Insert(8)
	fmt.Println(binarySearchTree.Find(8))
	binarySearchTree.Delete(6)
	fmt.Println(binarySearchTree.Find(7))
}
