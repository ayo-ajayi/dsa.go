package stack

type QueueWithStacks struct {
	stackIn  *Stack
	stackOut *Stack
}

func NewQueueWithStacks() *QueueWithStacks {
	return &QueueWithStacks{
		stackIn:  NewStack(),
		stackOut: NewStack(),
	}
}

func (q *QueueWithStacks) Add(item any) {
	q.stackIn.Push(item)
}

func (q *QueueWithStacks) Remove() any {
	if q.stackOut.IsEmpty() {
		for !q.stackIn.IsEmpty() {
			q.stackOut.Push(q.stackIn.Pop())
		}
	}
	return q.stackOut.Pop()
}

func (q *QueueWithStacks) Peek() any {
	if q.stackOut.IsEmpty() {
		for !q.stackIn.IsEmpty() {
			q.stackOut.Push(q.stackIn.Pop())
		}
	}
	return q.stackOut.Peek()
}

func (q *QueueWithStacks) IsEmpty() bool {
	return q.stackIn.IsEmpty() && q.stackOut.IsEmpty()
}
