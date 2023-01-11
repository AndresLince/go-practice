/**
 * 3.3
 * Stack of Plates: Imagine a (literal) stack of plates. If the stack gets
 * too high, it might topple. Therefore, in real life, we would likely
 * start a new stack when the previous stack exceeds some threshold.
 * Implement a data structure SetOfStacks that mimics this. SetOfStacks
 * should be composed of several stacks and should create a new stack once
 * the previous one exceeds capacity.
 * SetOfStacks. push() and SetOfStacks. pop() should behave identically to
 * a single stack (that is, pop () should return the same values as it
 * would if there were just a single stack).
 * FOLLOW UP
 * Implement a function popAt ( int index) which performs a pop operation
 * on a specific sub-stack.
 */
package stacks

import (
	"errors"
)

type SetOfStacks struct {
	arrayStacks  []MyStack
	length       int
	currentStack int
}

func NewSetOfStacks(length int) *SetOfStacks {
	stack1 := MyStack{}
	stack2 := MyStack{}
	stack3 := MyStack{}
	arrayStacks := []MyStack{
		stack1, stack2, stack3,
	}

	sos := SetOfStacks{
		arrayStacks:  arrayStacks,
		length:       length,
		currentStack: 0,
	}
	return &sos
}

func (setOfStacks *SetOfStacks) Push(data interface{}) error {
	if setOfStacks.currentStack == 2 && setOfStacks.arrayStacks[setOfStacks.currentStack].Length() == setOfStacks.length {
		return errors.New("full_stack")
	}
	if setOfStacks.currentStack < 2 && setOfStacks.arrayStacks[setOfStacks.currentStack].Length() == setOfStacks.length {
		setOfStacks.currentStack++
	}

	setOfStacks.arrayStacks[setOfStacks.currentStack].Push(data)
	return nil
}

func (setOfStacks *SetOfStacks) Pop() (interface{}, error) {
	if setOfStacks.currentStack == 0 && setOfStacks.arrayStacks[setOfStacks.currentStack].Length() == 0 {
		return nil, errors.New("empty_stack")
	}
	if setOfStacks.arrayStacks[setOfStacks.currentStack].IsEmpty() && setOfStacks.currentStack > 0 {
		setOfStacks.currentStack--
		return setOfStacks.Pop()
	}
	data := setOfStacks.arrayStacks[setOfStacks.currentStack].Pop()
	if setOfStacks.arrayStacks[setOfStacks.currentStack].IsEmpty() && setOfStacks.currentStack > 0 {
		setOfStacks.currentStack--
	}

	return data, nil
}

func (setOfStacks *SetOfStacks) IsEmpty() bool {
	return setOfStacks.arrayStacks[setOfStacks.currentStack].IsEmpty()
}

func (setOfStacks *SetOfStacks) Peak() interface{} {
	return setOfStacks.arrayStacks[setOfStacks.currentStack].Peek()
}

func (setOfStacks SetOfStacks) PopAt(stackNumber int) interface{} {
	if setOfStacks.arrayStacks[setOfStacks.currentStack].IsEmpty() {
		return nil
	}
	return setOfStacks.arrayStacks[stackNumber].Pop()
}
