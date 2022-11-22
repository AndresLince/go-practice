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

/**
 * 2.4
 * Partition: Write code to partition a linked list around a value x, such
 * that all nodes less than x come before all nodes greater than or equal
 * to x. If x is contained within the list, the values of x only need to be
 * after the elements less than x (see below). The partition element x can
 * appear anywhere in the "right partition"; it does not need to appear
 * between the left and right partitions.
 * EXAMPLE
 * Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition= 5]
 * Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
 */
func Partition(head Node, partition int) Node {
	node := head
	newLinkedList := Node{3, nil}
	for node.next != nil {
		if node.next.data < partition {
			newLinkedList = newLinkedList.Prepend(node.next.data)
		} else {
			newLinkedList.AppendToTail(node.next.data)
		}
		node = *node.next
	}
	return newLinkedList
}
