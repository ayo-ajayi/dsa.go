package stack

type Stack struct {
	top *StackNode
}

type StackNode struct {
	data any
	next *StackNode
}

func NewStackNode(d any) *StackNode {
	return &StackNode{data: d}
}
func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() any {
	if s.top == nil {
		return nil
	}
	item := s.top.data
	s.top = s.top.next
	return item
}


func (s *Stack) Push(item any) {
	newtop := NewStackNode(item)
	newtop.next = s.top
	s.top = newtop
}

func (s *Stack) Peek() any {
	if s.top == nil {
		return nil
	}
	return s.top.data
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}


// only works with int
func (s *Stack) Min() any {
	if s.top == nil {
		return nil
	}
	tempStack := s.top
	min := tempStack.data
	for tempStack.next != nil {
		if (tempStack.next.data).(int) < min.(int) {
			min = tempStack.next.data
		}
		tempStack = tempStack.next
	}
	return min
}
