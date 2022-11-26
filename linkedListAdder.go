package main

// LinkedListAdder known how to Add a number
type LinkedListAdder interface {
	// Concrete types should implement this method.
	Add(node *Node, number int)
	ResultAdd(node *Node, number int)
}

// LinkedListPrependAdder is a concrete type that implements Add
type LinkedListPrependAdder struct{}

// Add -- LinkedListPrependAdder adds the number
func (spm LinkedListPrependAdder) Add(node *Node, number int) {
	node.Prepend(number)
}

// ResultAdd -- LinkedListPrependAdder adds the number
func (sum LinkedListPrependAdder) ResultAdd(node *Node, number int) {
	node.AppendToTail(number)
}

// LinkedListAppendToTailAdder is a concrete type that implements Add
type LinkedListAppendToTailAdder struct{}

// Add -- LinkedListAppendToTailAdder adds the number
func (sum LinkedListAppendToTailAdder) Add(node *Node, number int) {
	node.AppendToTail(number)
}

// ResultAdd -- LinkedListAppendToTailAdder adds the number
func (sum LinkedListAppendToTailAdder) ResultAdd(node *Node, number int) {
	node.Prepend(number)
}

type linkedListStrategy struct {
	// LinkedListAdder is the behaviour that is encapsulated
	// Now this LinkedListAdder is of interface type and hence
	// can hold any concrete type.
	// Now the concrete type implements the methods of the
	// LinkedListAdder interface.
	LinkedListAdder LinkedListAdder
}

// NewLinkedListAdder returns a toy object
func NewLinkedListAdder(dr LinkedListAdder) *linkedListStrategy {
	return &linkedListStrategy{
		LinkedListAdder: dr,
	}
}

// AddToLinkedList performs the addition
func (t *linkedListStrategy) AddToLinkedList(node *Node, number int) {
	t.LinkedListAdder.Add(node, number)
}

// AddToLinkedList performs the addition to the result
func (t *linkedListStrategy) AddToLinkedListResult(node *Node, number int) {
	t.LinkedListAdder.ResultAdd(node, number)
}
