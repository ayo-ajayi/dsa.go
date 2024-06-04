package queue

// implement a Stack with two Queues
type StackWithQueues struct {
	queueIn *Queue
	buffer  *Queue
}

func NewStackWithQueues() *StackWithQueues {
	return &StackWithQueues{
		queueIn: NewQueue(),
		buffer:  NewQueue(),
	}
}

func (s *StackWithQueues) Push(item any) {
	s.buffer.Add(item)
	for !s.queueIn.IsEmpty() {
		s.buffer.Add(s.queueIn.Remove())
	}
	s.queueIn, s.buffer = s.buffer, s.queueIn
}

func (s *StackWithQueues) Peek() any {
	if s.queueIn.IsEmpty() {
		return nil
	}
	return s.queueIn.Peek()
}

func (s *StackWithQueues) IsEmpty() bool {
	return s.queueIn.IsEmpty()
}

func (s *StackWithQueues) Pop() any {
	if s.queueIn.IsEmpty() {
		return nil
	}
	return s.queueIn.Remove()
}
