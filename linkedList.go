package main

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
