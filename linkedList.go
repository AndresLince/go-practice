package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func (node *Node) AppendToTail(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}
	if node == nil {
		return
	}

	for node.next != nil {
		node = node.next
	}
	node.next = newNode
}

func (node *Node) Prepend(data int) {
	newNode := Node{}
	newNode.data = data

	if node == nil {
		*node = newNode
		return
	}
	head := node
	newNode.AppendToTail(node.data)
	for head.next != nil {
		newNode.AppendToTail(head.next.data)
		head = head.next
	}

	*node = newNode
}

func (node *Node) DeleteNode(data int) *Node {
	head := node
	if node.data == data {
		return head.next
	}
	for node.next != nil {
		if node.next.data == data {
			node.next = node.next.next
			return head
		}
		node = node.next
	}
	return head
}

func (node *Node) Print() {
	for node.next != nil {
		fmt.Println(node)
		node = node.next
	}
	fmt.Println(node)
}

func (node *Node) Length() int {
	if node == nil {
		return 0
	}
	counter := 1
	for node.next != nil {
		node = node.next
		counter++
	}
	return counter
}
