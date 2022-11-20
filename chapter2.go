package main

/**
2.2
Return Kth to Last: Implement an algorithm to find the kth to last
element of a singly linked list.
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

/**
2.3
Delete Middle Node: Implement an algorithm to delete a node in the middle
(i.e., any node but the first and last node, not necessarily the exact
middle) of a singly linked list, given only access to that node.
*/
func DeleteMiddleNode(head Node) *Node {
	node := head
	runner := head
	for runner.next != nil {
		runner = *runner.next.next
		node = *node.next
	}
	return head.DeleteNode(node.data)
}
