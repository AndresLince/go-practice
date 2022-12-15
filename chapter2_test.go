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
	expectedLinkedList := Node{5, &Node{7, &Node{4, nil}}}
	linkedListAdder := NewLinkedListAdder(LinkedListPrependAdder{})
	result := SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = Node{2, &Node{3, nil}}
	linkedList2 = Node{1, &Node{4, &Node{8, nil}}}
	expectedLinkedList = Node{3, &Node{7, &Node{8, nil}}}
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = Node{7, &Node{1, &Node{6, nil}}}
	linkedList2 = Node{5, &Node{9, &Node{2, nil}}}
	expectedLinkedList = Node{2, &Node{1, &Node{9, nil}}}
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = Node{9, &Node{9, &Node{9, nil}}}
	linkedList2 = Node{9, &Node{9, &Node{9, nil}}}
	expectedLinkedList = Node{8, &Node{9, &Node{9, &Node{1, nil}}}}
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}

	linkedListAdder = NewLinkedListAdder(LinkedListAppendToTailAdder{})
	linkedList1 = Node{4, &Node{3, &Node{1, nil}}}
	linkedList2 = Node{4, &Node{1, nil}}
	expectedLinkedList = Node{1, &Node{4, &Node{8, nil}}}
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`SumLists() expect valid linkedList`)
	}
	linkedList1 = Node{7, &Node{1, &Node{6, nil}}}
	linkedList2 = Node{5, &Node{9, &Node{2, nil}}}
	expectedLinkedList = Node{9, &Node{1, &Node{2, nil}}}
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`SumLists() expect valid linkedList`)
	}
}

func TestSumListsForwardOrder(t *testing.T) {
	linkedList1 := Node{1, &Node{3, &Node{4, nil}}}
	linkedList2 := Node{1, &Node{4, nil}}
	expectedLinkedList := Node{1, &Node{4, &Node{8, nil}}}
	result := SumListsForwardOrder(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumListsForwardOrder() expect valid linkedList`)
	}
	linkedList1 = Node{6, &Node{1, &Node{7, nil}}}
	linkedList2 = Node{2, &Node{9, &Node{5, nil}}}
	expectedLinkedList = Node{9, &Node{1, &Node{2, nil}}}
	result = SumListsForwardOrder(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumListsForwardOrder() expect valid linkedList`)
	}
}

func TestIsPalindrome(t *testing.T) {
	linkedList := &Node{"D", &Node{"a", &Node{"b", &Node{"a", &Node{"l",
		&Node{"e", &Node{"a", &Node{"r", &Node{"r", &Node{"o", &Node{"z",
			&Node{"a", &Node{"l", &Node{"a", &Node{"z", &Node{"o",
				&Node{"r", &Node{"r", &Node{"a", &Node{"e", &Node{"l",
					&Node{"a", &Node{"b", &Node{"a", &Node{"d", nil}}}}}}}}}}}}}}}}}}}}}}}}}
	result := IsPalindrome(linkedList)
	if result != true {
		t.Fatalf(`TestIsPalindrome() expect true`)
	}
	linkedList = &Node{"P", &Node{"e", &Node{"r", &Node{"r", &Node{"o", nil}}}}}
	result = IsPalindrome(linkedList)
	if result != false {
		t.Fatalf(`TestIsPalindrome() expect false`)
	}
}

func TestIsIntersection(t *testing.T) {
	intersectionNode := &Node{"d", &Node{"a", &Node{"s", nil}}}
	linkedList1 := &Node{"d", &Node{"a", &Node{"z", intersectionNode}}}
	linkedList2 := &Node{"p", &Node{"e", &Node{"r", &Node{"r", &Node{"o", intersectionNode}}}}}
	result := IsIntersection(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, intersectionNode) {
		t.Fatalf(`TestIsIntersection() expect a valid linked list`)
	}
	linkedList1 = &Node{"a", &Node{"b", &Node{"c", nil}}}
	linkedList2 = &Node{"z", &Node{"b", &Node{"c", nil}}}
	result = IsIntersection(linkedList1, linkedList2)
	if result != nil {
		t.Fatalf(`TestIsIntersection() expect nil`)
	}
}

func TestLoopDetection(t *testing.T) {
	linkedList := &Node{"d", &Node{"a", &Node{"z", nil}}}
	linkedList.next.next.next = linkedList
	result := LoopDetection(linkedList)
	if !reflect.DeepEqual(result.next, linkedList) {
		t.Fatalf(`TestLoopDetection() expect a valid linked list`)
	}
	tail := &Node{"d", &Node{"a", &Node{"z", nil}}}
	linkedList = &Node{"d", &Node{"a", &Node{"z", nil}}}
	linkedList.next.next.next = tail
	result = LoopDetection(linkedList)
	if result != nil {
		t.Fatalf(`TestIsIntersection() expect nil`)
	}
}
