package main

import "fmt"

type MyStack struct {
	top *Node
	min *Node
}

func (stack *MyStack) Push(data interface{}) {
	item := &Node{data, nil}
	item.next = stack.top
	stack.top = item
	if stack.min != nil && stack.top.data.(int) < stack.min.data.(int) {
		stack.min = item
	}
	if stack.min == nil {
		stack.min = stack.top
	}
}

func (stack *MyStack) Pop() interface{} {
	if stack.top == nil {
		return nil
	}
	data := stack.top.data
	if stack.min != nil && stack.min.data == data {
		stack.min = stack.top.GetMin()
	}
	stack.top = stack.top.next
	if stack.top == nil {
		stack.min = nil
	}

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

/**
 * 3.2
 * Stack Min: How would you design a stack which, in addition to push and
 * pop, has a function min which returns the minimum element? Push, pop and
 * min should all operate in 0(1) time.
 */
func (stack *MyStack) GetMin() *Node {
	fmt.Println("stack", stack)
	return stack.min
}
