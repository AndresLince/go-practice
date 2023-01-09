package stacks

import (
	"strings"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	stack := NewThreeInOneStack(6)
	if stack.IsEmpty(1) != true {
		t.Fatalf(`IsEmpty() expect true`)
	}
	stack.Push(1, 1)
	if stack.IsEmpty(1) != false {
		t.Fatalf(`IsEmpty() expect false`)
	}
}

func TestPeak(t *testing.T) {
	stack := NewThreeInOneStack(12)
	if stack.Peak(1) != nil {
		t.Fatalf(`TestPeak() expect nil`)
	}
	stack.Push(1, 1)
	if stack.Peak(1) != 1 {
		t.Fatalf(`TestPeak() expect 1`)
	}
	stack.Push(1, 2)
	if stack.Peak(1) != 2 {
		t.Fatalf(`TestPeak() expect 2`)
	}
	stack.Push(1, 3)
	if stack.Peak(1) != 3 {
		t.Fatalf(`TestPeak() expect 3`)
	}
}

func TestPop(t *testing.T) {
	stack := NewThreeInOneStack(12)
	stack.Push(1, 1)
	stack.Push(1, 2)
	stack.Push(1, 3)
	if stack.Peak(1) != 3 {
		t.Fatalf(`TestPeak() expect 3`)
	}
	stack.Pop(1)
	if stack.Peak(1) != 2 {
		t.Fatalf(`TestPeak() expect 2`)
	}
	stack.Pop(1)
	if stack.Peak(1) != 1 {
		t.Fatalf(`TestPeak() expect 2`)
	}
	stack.Pop(1)
	_, err := stack.Pop(1)
	if !ErrorContains(err, "empty_stack") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPush(t *testing.T) {
	stack := NewThreeInOneStack(12)
	err := stack.Push(1, 1)
	if ErrorContains(err, "full_stack") {
		t.Errorf("TestPeak() unexpected error: %v", err)
	}
	if stack.Peak(1) != 1 {
		t.Fatalf(`TestPeak() expect 1`)
	}
	err = stack.Push(1, 2)
	if ErrorContains(err, "full_stack") {
		t.Errorf("TestPeak() unexpected error: %v", err)
	}
	if stack.Peak(1) != 2 {
		t.Fatalf(`TestPeak() expect 2`)
	}
	err = stack.Push(1, 3)
	if ErrorContains(err, "full_stack") {
		t.Errorf("TestPeak() unexpected error: %v", err)
	}
	if stack.Peak(1) != 3 {
		t.Fatalf(`TestPeak() expect 3`)
	}
	err = stack.Push(1, 4)
	if ErrorContains(err, "full_stack") {
		t.Errorf("TestPeak() unexpected error: %v", err)
	}
	if stack.Peak(1) != 4 {
		t.Fatalf(`TestPeak() expect 4`)
	}

	err = stack.Push(1, 3)
	if !ErrorContains(err, "full_stack") {
		t.Errorf("TestPeak() unexpected error: %v", err)
	}
}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
