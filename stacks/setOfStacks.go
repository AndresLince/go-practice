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
