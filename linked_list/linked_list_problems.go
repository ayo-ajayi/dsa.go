package linked_list

// remove duplicates from an unsorted linked list
func (n *Node) RemoveDups() *Node {
	head := n
	store := map[any]int{}
	var prev *Node = nil

	for n != nil {
		if store[n.data] != 0 {
			prev.next = n.next
		} else {
			store[n.data]++
			prev = n
		}
		n = n.next

	}
	return head
}

func (n *Node) KthToLast(k int) any {
	store := map[int]any{}
	index := 0
	for n != nil {
		store[index] = n.data
		index++
		n = n.next
	}
	for i := 0; i < len(store); i++ {
		pos := len(store) - k
		if i == pos {
			return store[pos]
		}
	}
	return 0
}
