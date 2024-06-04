package tree

import (
	"container/list"
	"fmt"
)

// given a sorted (increasing order) array, create binary search tree with minmal height
func MinimalTree(sortedArr []int) *BinaryTreeNode {
	return minimalTree(sortedArr, 0, len(sortedArr)-1)
}

func minimalTree(sortedArr []int, start, end int) *BinaryTreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := NewBinaryTree(sortedArr[mid])
	node.left = minimalTree(sortedArr, start, mid-1)
	node.right = minimalTree(sortedArr, mid+1, end)
	return node
}

func (tree *BinaryTreeNode) MapOfDepths() {
	depth := 0
	track := map[int][]any{}
	tree.mapOfDepths(track, depth)
	fmt.Println(track)
}

func (tree *BinaryTreeNode) mapOfDepths(track map[int][]any, depth int) {
	if tree == nil {
		return
	}
	if track[depth] == nil {
		track[depth] = []any{tree.data}
	} else {
		track[depth] = append(track[depth], tree.data)
	}
	(depth)++
	tree.left.mapOfDepths(track, depth)
	tree.right.mapOfDepths(track, depth)
}

func (tree *BinaryTreeNode) LinkedListOfDepths() {
	l := list.New()
	tree.linkedListOfDepths(l, 0)

	for e, depth := l.Front(), 0; e != nil; e, depth = e.Next(), depth+1 {
		fmt.Print("Depth ", depth, ": ")
		ll := e.Value.(*list.List)
		for ele := ll.Front(); ele != nil; ele = ele.Next() {
			fmt.Print(ele.Value, " ")
		}
		fmt.Println()
	}
}

func (tree *BinaryTreeNode) linkedListOfDepths(l *list.List, depth int) {
	if tree == nil {
		return
	}

	if depth == l.Len() {
		l.PushBack(list.New())
	}

	depthList := l.Front()
	for i := 0; i < depth; i++ {
		depthList = depthList.Next()
	}
	depthList.Value.(*list.List).PushBack(tree.data)

	tree.left.linkedListOfDepths(l, depth+1)
	tree.right.linkedListOfDepths(l, depth+1)
}

func (tree *BinaryTreeNode) locateItemInTree(data int, depth int) (found bool, path string, depthFound int) {
	if tree == nil {
		return false, "", -1
	}
	if tree.data == data {
		return true, "root", depth
	}

	if tree.left != nil {
		depth++
		found, path, depthFound = tree.left.locateItemInTree(data, depth)
		if found {
			return true, "left -> " + path, depthFound
		}
	}

	if tree.right != nil {
		depth++
		found, path, depthFound = tree.right.locateItemInTree(data, depth)
		if found {
			return true, "right -> " + path, depthFound
		}
	}

	return false, "", -1
}

// LocateItemInTree starts the search for the item in the tree
func (tree *BinaryTreeNode) LocateItemInTree(data int) (depth int, path string) {
	found, path, depth := tree.locateItemInTree(data, 0)
	if !found {
		return -1, "not found" // Return -1 if the item is not found
	}
	return depth, path
}
