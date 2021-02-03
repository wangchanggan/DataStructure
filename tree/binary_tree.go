package tree

import "fmt"

type BinaryTree struct {
	Left  *BinaryTree
	Right *BinaryTree
	Data  interface{}
}

func NewBinaryTree(data interface{}) *BinaryTree {
	return new(BinaryTree).InitBinaryTree(data)
}

func (binaryTree *BinaryTree) InitBinaryTree(data interface{}) *BinaryTree {
	binaryTree.Data = data
	binaryTree.Left = nil
	binaryTree.Right = nil
	return binaryTree
}

/*
 * 前序遍历：对于树中的任意节点来说，先打印这个节点，然后再打印它的左子树，最后打印它的右子树。
 */
func (binaryTree *BinaryTree) PreOrder() {
	if binaryTree == nil {
		return
	}

	fmt.Println(binaryTree.Data)
	binaryTree.Left.PreOrder()
	binaryTree.Right.PreOrder()
}

/*
 * 中序遍历：对于树中的任意节点来说，先打印它的左子树，然后再打印它本身，最后打印它的右子树。
 */
func (binaryTree *BinaryTree) InOrder() {
	if binaryTree == nil {
		return
	}

	binaryTree.Left.InOrder()
	fmt.Println(binaryTree.Data)
	binaryTree.Right.InOrder()
}

/*
 * 后序遍历：对于树中的任意节点来说，先打印它的左子树，然后再打印它的右子树，最后打印这个节点本身。
 */
func (binaryTree *BinaryTree) PostOrder() {
	if binaryTree == nil {
		return
	}

	binaryTree.Left.PostOrder()
	binaryTree.Right.PostOrder()
	fmt.Println(binaryTree.Data)
}
