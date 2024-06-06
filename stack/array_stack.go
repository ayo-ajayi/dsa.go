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
		return nil 	//think of a better error handling technique
	}				//stack could contain (any) datatype that is nil...which would lead to confusion
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop

}

func (s *ArrayStack) Push(item any) {
	*s = append(*s, item)
}
