package main

/**
 * 2.2 Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
 */
func KthToLast(head Node, kth int) *Node {
	node := head
	runner := head
	counter := 0
	for runner.next != nil && counter < kth {
		runner = *runner.next
		counter++
	}
	if counter < kth {
		return nil
	}
	for runner.next != nil {
		node = *node.next
		runner = *runner.next
	}
	return &node
}
