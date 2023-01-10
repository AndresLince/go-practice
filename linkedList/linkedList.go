package linkedList

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

func (node *Node) AppendToTail(Data interface{}) {
	newNode := &Node{
		Data: Data,
		Next: nil,
	}

	for node.Next != nil {
		node = node.Next
	}
	node.Next = newNode
}

func (node *Node) Prepend(Data interface{}) {
	newNode := Node{}
	newNode.Data = Data

	head := node
	newNode.AppendToTail(node.Data)
	for head.Next != nil {
		newNode.AppendToTail(head.Next.Data)
		head = head.Next
	}

	*node = newNode
}

func (node *Node) DeleteNode(Data interface{}) *Node {
	head := node
	if node.Data == Data {
		return head.Next
	}
	for node.Next != nil {
		if node.Next.Data == Data {
			node.Next = node.Next.Next
			return head
		}
		node = node.Next
	}
	return head
}

func (node *Node) Print() {
	fmt.Println("--------------------")
	for node.Next != nil {
		fmt.Println(node)
		node = node.Next
	}
	fmt.Println(node)
}

func (node *Node) Length() int {
	counter := 1
	for node.Next != nil {
		node = node.Next
		counter++
	}
	return counter
}

func (node *Node) Revert() {
	newLinkedList := Node{}
	newLinkedList.Data = node.Data
	head := node
	for head.Next != nil {
		newLinkedList.Prepend(head.Next.Data)
		head = head.Next
	}
	*node = newLinkedList
}

func (node *Node) GetMin() *Node {
	head := node
	min := node
	for head != nil {
		if head.Data.(int) < min.Data.(int) {
			min = head
		}
		head = head.Next
	}
	return min
}
