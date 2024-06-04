package stack

type ArrayStack []any

func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

func (s *ArrayStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *ArrayStack) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	return (*s)[len(*s)-1]

}

func (s *ArrayStack) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop

}

func (s *ArrayStack) Push(item any) {
	*s = append(*s, item)
}
