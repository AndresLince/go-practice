package stacks

import (
	"fmt"

	"github.com/AndresLince/go-practice/linkedList"
)

type MyStack struct {
	top *linkedList.Node
	min *linkedList.Node
}

func (stack *MyStack) Push(data interface{}) {
	item := &linkedList.Node{data, nil}
	item.Next = stack.top
	stack.top = item
	if stack.min != nil && stack.top.Data.(int) < stack.min.Data.(int) {
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
	data := stack.top.Data
	stack.top = stack.top.Next
	if stack.top == nil {
		stack.min = nil
	}
	if stack.min != nil && stack.min.Data == data {
		stack.min = stack.top.GetMin()
	}

	return data
}

func (stack *MyStack) Peek() interface{} {
	if stack.top == nil {
		return nil
	}
	return stack.top.Data
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
func (stack *MyStack) GetMin() *linkedList.Node {
	fmt.Println("stack", stack)
	return stack.min
}
