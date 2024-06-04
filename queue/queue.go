package queue

type Queue struct {
	first *QueueNode
	last  *QueueNode
}

type QueueNode struct {
	data any
	next *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}
func NewQueueNode(d any) *QueueNode {
	return &QueueNode{
		data: d,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Add(item any) {
	newQueueNode := NewQueueNode(item)
	if q.last != nil {
		q.last.next = newQueueNode
	}
	q.last = newQueueNode
	if q.IsEmpty() {
		q.first = q.last
	}
}


func (q *Queue) Remove() (any) {
	if q.IsEmpty() {
		return nil
	}
	data := q.first.data
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	return data
}

func (q *Queue) Peek() (any) {
	if q.IsEmpty() {
		return nil
	}
	return q.first.data
}
