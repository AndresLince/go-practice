package main

import (
	"fmt"
)

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
	newLinkedList := Node{node.data, nil}
	for node.next != nil {
		if node.next.data < partition {
			newLinkedList.Prepend(node.next.data)
		} else {
			newLinkedList.AppendToTail(node.next.data)
		}
		node = *node.next
	}
	return newLinkedList
}

/**
 * 2.5
 * Sum Lists: You have two numbers represented by a linked list, where each node contains a single
 * digit. The digits are stored in reverse order, such that the 1 's digit is at the head of the list. Write a
 * function that adds the two numbers and returns the sum as a linked list.
 * EXAMPLE
 * Input:(7-> 1 -> 6) + (5 -> 9 -> 2).That is,617 + 295.
 * Output: 2 -> 1 -> 9. That is, 912.
 * FOLLOW UP
 * Suppose the digits are stored in forward order. Repeat the above problem.
 * EXAMPLE
 * lnput:(6 -> 1 -> 7) + (2 -> 9 -> 5).That is,617 + 295.
 * Output: 9 -> 1 -> 2. That is, 912.
 */
func SumLists(list1 *Node, list2 *Node) Node {
	var s []int
	carry := 0

	for list1 != nil && list2 != nil {
		fmt.Println(list1.data, list2.data)

		result := list1.data + list2.data + carry
		carryover := result % 10
		if carryover > 0 && result > 10 {
			s = append(s, carryover)
			carry = 1
		} else {
			s = append(s, result)
			carry = 0
		}
		fmt.Println("carryover", carryover)
		list1 = list1.next
		list2 = list2.next
	}
	if carry > 0 {
		s = append(s, carry)
	}
	fmt.Println(s, carry)
	return ArrayToList(s)
}

func ArrayToList(array []int) Node {
	newLinkedList := Node{array[0], nil}
	for i, v := range array {
		if i == 0 {
			continue
		}
		newLinkedList.AppendToTail(v)
	}
	return newLinkedList
}
