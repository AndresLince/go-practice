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
}

func NewQueueWithStack() *QueueWithStack {
	stack1 := stacks.MyStack{}
	stack2 := stacks.MyStack{}
	qws := QueueWithStack{stack1: stack1, stack2: stack2}
	return &qws
}

func (queueWithStack *QueueWithStack) Add(data interface{}) {
	queueWithStack.stack1.Push(data)
}

func (queueWithStack QueueWithStack) Remove() interface{} {
	for !queueWithStack.stack1.IsEmpty() {
		queueWithStack.stack2.Push(queueWithStack.stack1.Pop())
	}
	data := queueWithStack.stack2.Pop()
	for !queueWithStack.stack2.IsEmpty() {
		queueWithStack.stack1.Push(queueWithStack.stack2.Pop())
	}
	return data
}
