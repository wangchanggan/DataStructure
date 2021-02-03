package tree

type BinarySearchTree struct {
	Left  *BinarySearchTree
	Right *BinarySearchTree
	Data  int
}

func NewBinarySearchTree(data int) *BinarySearchTree {
	return new(BinarySearchTree).InitBinarySearchTree(data)
}

func (binarySearchTree *BinarySearchTree) InitBinarySearchTree(data int) *BinarySearchTree {
	binarySearchTree.Data = data
	binarySearchTree.Left = nil
	binarySearchTree.Right = nil
	return binarySearchTree
}

func (binarySearchTree *BinarySearchTree) Find(data int) *BinarySearchTree {
	if binarySearchTree == nil {
		return nil
	}

	if binarySearchTree.Data == data {
		return binarySearchTree
	} else if data < binarySearchTree.Data {
		return binarySearchTree.Left.Find(data)
	} else {
		return binarySearchTree.Right.Find(data)
	}
}

func (binarySearchTree *BinarySearchTree) Insert(data int) {
	if binarySearchTree == nil {
		binarySearchTree = NewBinarySearchTree(data)
		return
	}

	bst := binarySearchTree
	for bst != nil {
		if data < bst.Data {
			if bst.Left == nil {
				bst.Left = NewBinarySearchTree(data)
				return
			}
			bst = bst.Left
		} else {
			if bst.Right == nil {
				bst.Right = NewBinarySearchTree(data)
				return
			}
			bst = bst.Right
		}
	}
}

func (binarySearchTree *BinarySearchTree) Delete(data int) {
	bst := binarySearchTree
	fbst := new(BinarySearchTree)

	if bst == nil {
		return
	}

	//记录要删除节点的父节点，并赋予fbst
	for bst != nil && bst.Data != data {
		fbst = bst
		if data < bst.Data {
			bst = bst.Left
		} else {
			bst = bst.Right
		}
	}

	//如果要删除的节点有两个子节点
	if bst.Left != nil && bst.Right != nil {
		//查找右子树中最小的节点
		minBst := bst.Right
		minFbst := bst //表示右子树中最小的节点的父节点
		for minBst.Left != nil {
			minFbst = minBst
			minBst = minBst.Left
		}
		bst.Data = minBst.Data
		bst = minBst
		fbst = minFbst
	}

	//删除节点是叶子节点或仅有一个子节点
	child := new(BinarySearchTree)
	if bst.Left != nil {
		child = bst.Left
	} else if bst.Right != nil {
		child = bst.Right
	} else {
		child = nil
	}

	if fbst == nil { //删除的是根结点
		binarySearchTree = child
	} else if fbst.Left == bst {
		fbst.Left = child
	} else {
		fbst.Right = child
	}
}
