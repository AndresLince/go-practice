package main

import (
	"reflect"
	"testing"
)

func TestKthToLast(t *testing.T) {
	linkedList := Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	expectedLinkedList := &Node{5, &Node{6, &Node{7, nil}}}
	response := KthToLast(linkedList, 2)
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

func TestPartition(t *testing.T) {
	linkedList := Node{3, &Node{5, &Node{8, &Node{5, &Node{10, &Node{2, &Node{1, nil}}}}}}}
	expectedLinkedList := Node{1, &Node{2, &Node{3, &Node{5, &Node{8, &Node{5, &Node{10, nil}}}}}}}
	result := Partition(linkedList, 5)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`Partition() expect valid linkedList`)
	}
}

func TestSumLists(t *testing.T) {
	linkedList1 := Node{4, &Node{3, &Node{4, nil}}}
	linkedList2 := Node{1, &Node{4, nil}}
	expectedLinkedList := Node{4, &Node{4, &Node{8, nil}}}
	result := SumLists(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = Node{7, &Node{1, &Node{6, nil}}}
	linkedList2 = Node{5, &Node{9, &Node{2, nil}}}
	expectedLinkedList = Node{2, &Node{1, &Node{9, nil}}}
	result = SumLists(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
}
