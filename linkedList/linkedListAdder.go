package linkedList

// LinkedListAdder known how to Add a number
type LinkedListAdder interface {
	// Concrete types should implement this method.
	ResultAdd(node *Node, number int)
}

// LinkedListPrependAdder is a concrete type that implements Add
type LinkedListPrependAdder struct{}

// ResultAdd -- LinkedListPrependAdder adds the number
func (sum LinkedListPrependAdder) ResultAdd(node *Node, number int) {
	node.AppendToTail(number)
}

// LinkedListAppendToTailAdder is a concrete type that implements Add
type LinkedListAppendToTailAdder struct{}

// ResultAdd -- LinkedListAppendToTailAdder adds the number
func (sum LinkedListAppendToTailAdder) ResultAdd(node *Node, number int) {
	node.Prepend(number)
}

type LinkedListStrategy struct {
	// LinkedListAdder is the behaviour that is encapsulated
	// Now this LinkedListAdder is of interface type and hence
	// can hold any concrete type.
	// Now the concrete type implements the methods of the
	// LinkedListAdder interface.
	LinkedListAdder LinkedListAdder
}

// NewLinkedListAdder returns a toy object
func NewLinkedListAdder(dr LinkedListAdder) *LinkedListStrategy {
	return &LinkedListStrategy{
		LinkedListAdder: dr,
	}
}

// AddToLinkedList performs the addition to the result
func (t *LinkedListStrategy) AddToLinkedListResult(node *Node, number int) {
	t.LinkedListAdder.ResultAdd(node, number)
}
