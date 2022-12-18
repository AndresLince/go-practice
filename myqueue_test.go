package main

import "testing"

func TestIsEmptyQueue(t *testing.T) {
	queue := MyQueue{}
	if queue.IsEmpty() != true {
		t.Fatalf(`IsEmpty() expect true`)
	}
	queue.Push(1)
	if queue.IsEmpty() != false {
		t.Fatalf(`IsEmpty() expect false`)
	}
}

func TestPeekQueue(t *testing.T) {
	queue := MyQueue{}
	if queue.Peek() != nil {
		t.Fatalf(`Peek() expect nil`)
	}
	queue.Push(1)
	if queue.Peek() != 1 {
		t.Fatalf(`Peek() expect 1`)
	}
	queue.Push(2)
	if queue.Peek() != 1 {
		t.Fatalf(`Peek() expect 1`)
	}
	queue.Push(3)
	if queue.Peek() != 1 {
		t.Fatalf(`Peek() expect 1`)
	}
}

func TestPopQueue(t *testing.T) {
	queue := MyQueue{}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if queue.Pop() != 1 {
		t.Fatalf(`Pop() expect 1`)
	}
	if queue.Pop() != 2 {
		t.Fatalf(`Pop() expect 1`)
	}
	if queue.Pop() != 3 {
		t.Fatalf(`Pop() expect 1`)
	}
	if queue.Pop() != nil {
		t.Fatalf(`Pop() expect 1`)
	}
}

func TestPushQueue(t *testing.T) {
	queue := MyQueue{}
	queue.Push(1)
	if queue.Peek() != 1 {
		t.Fatalf(`Push() expect 1`)
	}
	queue.Push(2)
	if queue.Peek() != 1 {
		t.Fatalf(`Push() expect 1`)
	}
	queue.Push(3)
	if queue.Peek() != 1 {
		t.Fatalf(`Push() expect 1`)
	}
}
