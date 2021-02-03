package tree

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	binaryTree := NewBinaryTree("A")
	binaryTree.Left = NewBinaryTree("B")
	binaryTree.Right = NewBinaryTree("C")
	binaryTree.Left.Left = NewBinaryTree("D")
	binaryTree.Left.Right = NewBinaryTree("E")
	binaryTree.Right.Left = NewBinaryTree("F")
	binaryTree.Right.Right = NewBinaryTree("G")
	fmt.Println("前序遍历：")
	binaryTree.PreOrder()
	fmt.Println("中序遍历：")
	binaryTree.InOrder()
	fmt.Println("后序遍历：")
	binaryTree.PostOrder()
}
