package main

type MyQueueNode struct {
	data interface{}
	next *MyQueueNode
}

type MyQueue struct {
	first *MyQueueNode
	last  *MyQueueNode
}

func (queue *MyQueue) Push(data interface{}) {
	node := &MyQueueNode{data, nil}
	if queue.last != nil {
		queue.last.next = node
	}
	queue.last = node
	if queue.first == nil {
		queue.first = node
	}
}

func (queue *MyQueue) Pop() interface{} {
	if queue.first == nil {
		return nil
	}
	data := queue.first.data
	queue.first = queue.first.next
	if queue.first == nil {
		queue.last = nil
	}
	return data
}

func (queue *MyQueue) Peek() interface{} {
	if queue.first == nil {
		return nil
	}
	return queue.first.data
}

func (queue *MyQueue) IsEmpty() bool {
	return queue.first == nil
}
