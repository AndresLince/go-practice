package main

import "testing"

func TestIsEmpty(t *testing.T) {
	stack := MyStack{}
	if !stack.IsEmpty() {
		t.Fatalf(`IsEmpty() expect true`)
	}
	stack.Push(1)
	if stack.IsEmpty() {
		t.Fatalf(`IsEmpty() expect false`)
	}
}

func TestPeek(t *testing.T) {
	stack := MyStack{}
	if stack.Peek() != nil {
		t.Fatalf(`TestPeek() expect nil`)
	}
	stack.Push(1)
	if stack.Peek() != 1 {
		t.Fatalf(`TestPeek() expect 1`)
	}
	stack.Push(2)
	if stack.Peek() != 2 {
		t.Fatalf(`TestPeek() expect 2`)
	}
	stack.Push(3)
	if stack.Peek() != 3 {
		t.Fatalf(`TestPeek() expect 3`)
	}
}

func TestPop(t *testing.T) {
	stack := MyStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Peek() != 3 {
		t.Fatalf(`TestPeek() expect 3`)
	}
	stack.Pop()
	if stack.Peek() != 2 {
		t.Fatalf(`TestPeek() expect 2`)
	}
	stack.Pop()
	if stack.Peek() != 1 {
		t.Fatalf(`TestPeek() expect 2`)
	}
	stack.Pop()
	if stack.Peek() != nil {
		t.Fatalf(`TestPeek() expect nil`)
	}
	stack.Pop()
	if stack.Peek() != nil {
		t.Fatalf(`TestPeek() expect nil`)
	}
}
func TestPush(t *testing.T) {
	stack := MyStack{}
	stack.Push(1)
	if stack.Peek() != 1 {
		t.Fatalf(`TestPeek() expect 1`)
	}
	stack.Push(2)
	if stack.Peek() != 2 {
		t.Fatalf(`TestPeek() expect 2`)
	}
	stack.Push(3)
	if stack.Peek() != 3 {
		t.Fatalf(`TestPeek() expect 3`)
	}
}

func TestGetMinStack(t *testing.T) {
	stack := MyStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.Push(4)
	if stack.GetMin().data != 1 {
		t.Fatalf(`GetMin() expect 1`)
	}
	stack = MyStack{}
	stack.Push(3)
	stack.Push(4)
	stack.Push(1)
	stack.Pop()
	if stack.GetMin().data != 3 {
		t.Fatalf(`GetMin() expect 3`)
	}
}
