package main

import "fmt"

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
