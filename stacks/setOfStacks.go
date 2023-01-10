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

type SetOfStacks struct {
	stack1 MyStack
	stack2 MyStack
	stack3 MyStack
	length int
}

func NewSetOfStacks(length int) *SetOfStacks {
	stack1 := MyStack{}
	stack2 := MyStack{}
	stack3 := MyStack{}

	sos := SetOfStacks{
		stack1: stack1,
		stack2: stack2,
		stack3: stack3,
		length: length,
	}
	return &sos
}
