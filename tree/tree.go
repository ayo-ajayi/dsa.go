package tree

import (
	"fmt"
)

type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func NewBinaryTree(data int) *BinaryTreeNode {
	return &BinaryTreeNode{
		data: data,
	}
}

func (n *BinaryTreeNode) Insert(data int) {

	if data < n.data {
		if n.left == nil {
			n.left = &BinaryTreeNode{data: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{data: data}
		} else {
			n.right.Insert(data)
		}
	}
}

func (n *BinaryTreeNode) InOrderTraversal() {
	if n != nil {
		n.left.InOrderTraversal()
		fmt.Print(n.data, " ")
		n.right.InOrderTraversal()
	}
}

func (n *BinaryTreeNode) PreOrderTraversal() {
	if n != nil {
		fmt.Print(n.data, " ")
		n.left.PreOrderTraversal()
		n.right.PreOrderTraversal()
	}

}

func (n *BinaryTreeNode) PostOrderTraversal() {
	if n != nil {
		n.left.PostOrderTraversal()
		n.right.PostOrderTraversal()
		fmt.Print(n.data, " ")
	}
}
