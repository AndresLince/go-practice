package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKthToLast(t *testing.T) {
	linkedList := Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	expectedLinkedList := &Node{5, &Node{6, &Node{7, nil}}}
	response := KthToLast(linkedList, 2)
	fmt.Println(response)
	fmt.Println(expectedLinkedList)
	if !reflect.DeepEqual(response, expectedLinkedList) {
		t.Fatalf(`KthToLast(linkedList, 2) expect a valid linked list`)
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	linkedList := Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	expectedLinkedList := &Node{1, &Node{2, &Node{3, &Node{5, &Node{6, &Node{7, nil}}}}}}
	result := DeleteMiddleNode(linkedList)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`DeleteMiddleNode() expect valid linkedList`)
	}
}
