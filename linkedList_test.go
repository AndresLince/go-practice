package main

import (
	"reflect"
	"testing"
)

func TestAppendToTail(t *testing.T) {
	expectedLinkedList := Node{1, &Node{2, &Node{3, nil}}}
	linkedList := Node{1, nil}
	linkedList.AppendToTail(2)
	linkedList.AppendToTail(3)
	if !reflect.DeepEqual(linkedList, expectedLinkedList) {
		t.Fatalf(`AppendToTail() expect valid linkedList`)
	}
}

func TestPrepend(t *testing.T) {
	linkedList := Node{2, &Node{3, nil}}
	expectedLinkedList := Node{1, &Node{2, &Node{3, nil}}}
	linkedList.Prepend(1)
	if !reflect.DeepEqual(linkedList, expectedLinkedList) {
		t.Fatalf(`Prepend() expect valid linkedList`)
	}
}

func TestDeleteNode(t *testing.T) {
	linkedList := Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	expectedLinkedList := &Node{1, &Node{2, &Node{3, &Node{5, &Node{6, &Node{7, nil}}}}}}
	result := linkedList.DeleteNode(4)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`DeleteNode() expect valid linkedList`)
	}
	linkedList = Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	expectedLinkedList = &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}
	result = linkedList.DeleteNode(1)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`DeleteNode() expect valid linkedList`)
	}
	linkedList = Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	expectedLinkedList = &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, nil}}}}}}
	result = linkedList.DeleteNode(7)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`DeleteNode() expect valid linkedList`)
	}
	linkedList = Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	result = linkedList.DeleteNode(7)
	if !reflect.DeepEqual(result, &linkedList) {
		t.Fatalf(`DeleteNode() expect valid linkedList`)
	}
	linkedList = Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	result = linkedList.DeleteNode(10)
	if !reflect.DeepEqual(result, &linkedList) {
		t.Fatalf(`DeleteNode() expect valid linkedList`)
	}
}

func TestLength(t *testing.T) {
	linkedList := Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	result := linkedList.Length()
	if result != 7 {
		t.Fatalf(`Length() expect 7`)
	}
	linkedList = Node{}
	result = linkedList.Length()
	if result != 1 {
		t.Fatalf(`Length() expect 1`)
	}
	linkedList = Node{5, &Node{4, &Node{3, &Node{2, nil}}}}
	result = linkedList.Length()
	if result != 4 {
		t.Fatalf(`Length() expect 4`)
	}
}

func TestRevert(t *testing.T) {
	linkedList := Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	expectedLinkedList := Node{7, &Node{6, &Node{5, &Node{4, &Node{3, &Node{2, &Node{1, nil}}}}}}}
	linkedList.Revert()
	if !reflect.DeepEqual(linkedList, expectedLinkedList) {
		t.Fatalf(`Revert() expect a valid linked list`)
	}
	linkedList = Node{3, &Node{2, &Node{1, nil}}}
	expectedLinkedList = Node{1, &Node{2, &Node{3, nil}}}
	linkedList.Revert()
	if !reflect.DeepEqual(linkedList, expectedLinkedList) {
		t.Fatalf(`Revert() expect a valid linked list`)
	}
}

func TestGetMin(t *testing.T) {
	linkedList := Node{5, &Node{4, &Node{3, &Node{6, &Node{7, &Node{1, &Node{2, nil}}}}}}}
	response := linkedList.GetMin()
	if response.data != 1 {
		t.Fatalf(`GetMin().data expect a valid value`)
	}
}
