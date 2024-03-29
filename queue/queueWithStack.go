/**
 * 3.4
 * Queue via Stacks: Implement a MyQueue class which implements a queue using
 * two stacks
 */
package queue

import (
	"github.com/AndresLince/go-practice/stacks"
)

type QueueWithStack struct {
	stack1 stacks.MyStack
	stack2 stacks.MyStack
	head   interface{}
}

func NewQueueWithStack() *QueueWithStack {
	stack1 := stacks.MyStack{}
	stack2 := stacks.MyStack{}
	qws := QueueWithStack{stack1: stack1, stack2: stack2}
	return &qws
}

func (queueWithStack *QueueWithStack) Add(data interface{}) {
	if queueWithStack.stack1.IsEmpty() {
		queueWithStack.head = data
	}
	queueWithStack.stack1.Push(data)
}

func (queueWithStack *QueueWithStack) Remove() interface{} {
	for !queueWithStack.stack1.IsEmpty() {
		queueWithStack.stack2.Push(queueWithStack.stack1.Pop())
	}
	data := queueWithStack.stack2.Pop()
	if queueWithStack.stack2.IsEmpty() {
		queueWithStack.head = nil
	} else {
		queueWithStack.head = queueWithStack.stack2.Peek()
	}
	for !queueWithStack.stack2.IsEmpty() {
		queueWithStack.stack1.Push(queueWithStack.stack2.Pop())
	}
	return data
}

func (queueWithStack QueueWithStack) Peak() interface{} {
	return queueWithStack.head
}

func (queueWithStack QueueWithStack) IsEmpty() bool {
	return queueWithStack.stack1.IsEmpty()
}
