package stacks

import (
	"testing"
)

func TestIsEmptySetOfStacks(t *testing.T) {
	stack := NewSetOfStacks(3)
	if stack.IsEmpty() != true {
		t.Fatalf(`IsEmpty() expect true`)
	}
	stack.Push(1)
	if stack.IsEmpty() != false {
		t.Fatalf(`IsEmpty() expect false`)
	}
}

func TestPeakSetOfStacks(t *testing.T) {
	stack := NewSetOfStacks(3)
	if stack.Peak() != nil {
		t.Fatalf(`Peak() expect nil`)
	}
	stack.Push(1)
	if stack.Peak() != 1 {
		t.Fatalf(`Peak() expect 1`)
	}
	stack.Push(2)
	if stack.Peak() != 2 {
		t.Fatalf(`Peak() expect 2`)
	}
	stack.Push(3)
	if stack.Peak() != 3 {
		t.Fatalf(`Peak() expect 3`)
	}
	stack.Push(4)
	if stack.Peak() != 4 {
		t.Fatalf(`Peak() expect 4`)
	}
	stack.Push(5)
	if stack.Peak() != 5 {
		t.Fatalf(`Peak() expect 5`)
	}
}

func TestPopSetOfStacks(t *testing.T) {
	stack := NewSetOfStacks(2)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	if stack.Peak() != 6 {
		t.Fatalf(`Peak() expect 6`)
	}
	_, err := stack.Pop()
	if ErrorContains(err, "empty_stack") {
		t.Errorf("TestPopSetOfStacks unexpected error: %v", err)
	}
	if stack.Peak() != 5 {
		t.Fatalf(`Peak() expect 5`)
	}
	_, err = stack.Pop()
	if ErrorContains(err, "empty_stack") {
		t.Errorf("TestPopSetOfStacks unexpected error: %v", err)
	}
	if stack.Peak() != 4 {
		t.Fatalf(`Peak() expect 4`)
	}
	_, err = stack.Pop()
	if ErrorContains(err, "empty_stack") {
		t.Errorf("TestPopSetOfStacks unexpected error: %v", err)
	}
	if stack.Peak() != 3 {
		t.Fatalf(`Peak() expect 3`)
	}
	_, err = stack.Pop()
	if ErrorContains(err, "empty_stack") {
		t.Errorf("TestPopSetOfStacks unexpected error: %v", err)
	}
	if stack.Peak() != 2 {
		t.Fatalf(`Peak() expect 2`)
	}
	_, err = stack.Pop()
	if ErrorContains(err, "empty_stack") {
		t.Errorf("TestPopSetOfStacks unexpected error: %v", err)
	}
	if stack.Peak() != 1 {
		t.Fatalf(`Peak() expect 2`)
	}
	_, err = stack.Pop()
	if ErrorContains(err, "empty_stack") {
		t.Errorf("TestPopSetOfStacks unexpected error: %v", err)
	}
	_, err = stack.Pop()
	if !ErrorContains(err, "empty_stack") {
		t.Errorf("TestPopSetOfStacks expected error: %v", err)
	}
}
