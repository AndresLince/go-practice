package main

type MyStack struct {
	top *Node
}

func (stack *MyStack) Push(data interface{}) {
	item := &Node{data, nil}
	item.next = stack.top
	stack.top = item
}

func (stack *MyStack) Pop() interface{} {
	if stack.top == nil {
		return nil
	}
	data := stack.top.data
	stack.top = stack.top.next
	return data
}

func (stack *MyStack) Peek() interface{} {
	if stack.top == nil {
		return nil
	}
	return stack.top.data
}

func (stack *MyStack) IsEmpty() bool {
	return stack.top == nil
}
