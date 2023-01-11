package main

import (
	"math"
	"strings"

	"github.com/AndresLince/go-practice/linkedList"
)

/**
2.2
Return Kth to Last: Implement an algorithm to find the kth to last
element of a singly linked list.
*/
func KthToLast(head linkedList.Node, kth int) *linkedList.Node {
	node := head
	runner := head
	counter := 0
	for runner.Next != nil && counter < kth {
		runner = *runner.Next
		counter++
	}
	if counter < kth {
		return nil
	}
	for runner.Next != nil {
		node = *node.Next
		runner = *runner.Next
	}
	return &node
}

/**
2.3
Delete Middle Node: Implement an algorithm to delete a node in the middle
(i.e., any node but the first and last node, not necessarily the exact
middle) of a singly linked list, given only access to that node.
*/
func DeleteMiddleNode(head linkedList.Node) *linkedList.Node {
	node := head
	runner := head
	for runner.Next != nil {
		runner = *runner.Next.Next
		node = *node.Next
	}
	return head.DeleteNode(node.Data)
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
func Partition(head linkedList.Node, partition int) linkedList.Node {
	node := head
	newLinkedList := linkedList.Node{Data: node.Data, Next: nil}
	for node.Next != nil {
		if node.Next.Data.(int) < partition {
			newLinkedList.Prepend(node.Next.Data)
		} else {
			newLinkedList.AppendToTail(node.Next.Data)
		}
		node = *node.Next
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
func SumLists(list1 linkedList.Node, list2 linkedList.Node, linkedListStrategy linkedList.LinkedListStrategy) linkedList.Node {
	newList1 := &list1
	newList2 := &list2

	AddZeros(newList1, newList2)

	var s []int
	carry := 0

	for newList1 != nil && newList2 != nil {

		result := newList1.Data.(int) + newList2.Data.(int) + carry
		carryover := result % 10
		if carryover > 0 && result > 10 {
			s = append(s, carryover)
			carry = 1
		} else {
			s = append(s, result)
			carry = 0
		}

		newList1 = newList1.Next
		newList2 = newList2.Next
	}
	if carry > 0 {
		s = append(s, carry)
	}

	return ArrayToList(s, linkedListStrategy)
}

func ArrayToList(array []int, linkedListStrategy linkedList.LinkedListStrategy) linkedList.Node {
	newLinkedList := linkedList.Node{Data: array[0], Next: nil}
	for i, v := range array {
		if i == 0 {
			continue
		}
		linkedListStrategy.AddToLinkedListResult(&newLinkedList, v)
	}
	return newLinkedList
}

func AddZeros(list1 *linkedList.Node, list2 *linkedList.Node) {
	length1 := list1.Length()
	length2 := list2.Length()

	difference := math.Abs(float64((length1 - length2)))
	if difference > 0 && length1 > length2 {
		AddNumberToLinkedList(list2, int(difference), 0)
	}
	if difference > 0 && length1 < length2 {
		AddNumberToLinkedList(list1, int(difference), 0)
	}
}
func AddNumberToLinkedList(list1 *linkedList.Node, numberOfZeros int, number int) {
	for i := 0; i < int(numberOfZeros); i++ {
		list1.AppendToTail(number)
	}
}

func SumListsForwardOrder(list1 linkedList.Node, list2 linkedList.Node) linkedList.Node {
	list1.Revert()
	list2.Revert()

	NewLinkedListAdder := linkedList.NewLinkedListAdder(linkedList.LinkedListPrependAdder{})

	return SumLists(list1, list2, *NewLinkedListAdder)
}

/**
 * 2.6
 * Palindrome: Implement a function to check if a linked list is a
 * palindrome
 */

func IsPalindrome(list *linkedList.Node) bool {
	node := &linkedList.Node{Data: list.Data, Next: list.Next}
	list.Revert()
	for list != nil {
		if !strings.EqualFold(strings.ToLower(node.Data.(string)), strings.ToLower(list.Data.(string))) {
			return false
		}
		list = list.Next
		node = node.Next
	}

	return true
}

/**
 * 2.7
 * Intersection: Given two (singly) linked lists, determine if the two
 * lists intersect. Return the intersecting node. Note that the
 * intersection is defined based on reference, not value. That is, if the
 * kth node of the first linked list is the exact same node (by reference)
 * as the jth node of the second linked list, then they are intersecting.
 */
func IsIntersection(list1, list2 *linkedList.Node) *linkedList.Node {
	len2 := list2.Length()
	for list1 != nil {
		pointer := list2
		for i := 0; i < len2; i++ {
			if pointer == list1 {
				return list1
			}
			pointer = pointer.Next
		}
		list1 = list1.Next
	}
	return nil
}

/**
 * 2.8
 * Loop Detection: Given a circular linked list, implement an algorithm
 * that returns the node at the beginning of the loop.
 * DEFINITION
 * Circular linked list: A (corrupt) linked list in which a node's Next
 * pointer points to an earlier node, so as to make a loop in the linked
 * list.
 * EXAMPLE
 * Input: A -> B -> C -> D -> E -> C [the same C as earlier]
 * Output: C
 */
func LoopDetection(linkedList *linkedList.Node) *linkedList.Node {
	if linkedList == nil {
		return nil
	}
	head := linkedList
	for linkedList.Next != nil {
		if linkedList.Next == head {
			return linkedList
		}
		linkedList = linkedList.Next
	}
	return nil
}

/**
 * ----------------------------------------------------------------------------
 */
func CreateIncrementalNumberLinkedList(start int, end int) linkedList.Node {
	if start > end {
		return linkedList.Node{}
	}
	node := linkedList.Node{Data: start, Next: nil}
	start++
	for start <= end {
		node.AppendToTail(start)
		start++
	}
	return node
}
