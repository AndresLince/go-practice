package main

import (
	"reflect"
	"testing"

	"github.com/AndresLince/go-practice/linkedList"
)

func TestKthToLast(t *testing.T) {
	newLinkedList := CreateIncrementalNumberLinkedList(1, 7)
	expectedLinkedList := CreateIncrementalNumberLinkedList(5, 7)
	response := KthToLast(newLinkedList, 2)
	if !reflect.DeepEqual(response, &expectedLinkedList) {
		t.Fatalf(`KthToLast(linkedList, 2) expect a valid linked list`)
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	newLinkedList := CreateIncrementalNumberLinkedList(1, 7)
	expectedLinkedList := CreateIncrementalNumberLinkedList(1, 3)
	expectedLinkedList.AppendToTail(5)
	expectedLinkedList.AppendToTail(6)
	expectedLinkedList.AppendToTail(7)
	result := DeleteMiddleNode(newLinkedList)
	if !reflect.DeepEqual(result, &expectedLinkedList) {
		t.Fatalf(`DeleteMiddleNode() expect valid linkedList`)
	}
}

func TestPartition(t *testing.T) {
	newLinkedList := ArrayToList([]int{3, 5, 8, 5, 10, 2, 1}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList := ArrayToList([]int{1, 2, 3, 5, 8, 5, 10}, *NewLinkedListAppendToTailAdder())
	result := Partition(newLinkedList, 5)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`Partition() expect valid linkedList`)
	}
}

func TestSumLists(t *testing.T) {
	linkedList1 := ArrayToList([]int{4, 3, 4}, *NewLinkedListAppendToTailAdder())
	linkedList2 := ArrayToList([]int{1, 4}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList := ArrayToList([]int{5, 7, 4}, *NewLinkedListAppendToTailAdder())
	linkedListAdder := linkedList.NewLinkedListAdder(linkedList.LinkedListAppendToTailAdder{})
	result := SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]int{2, 3}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]int{1, 4, 8}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]int{3, 7, 8}, *NewLinkedListAppendToTailAdder())
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]int{7, 1, 6}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]int{5, 9, 2}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]int{2, 1, 9}, *NewLinkedListAppendToTailAdder())
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]int{9, 9, 9}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]int{9, 9, 9}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]int{8, 9, 9, 1}, *NewLinkedListAppendToTailAdder())
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}

	linkedListAdder = linkedList.NewLinkedListAdder(linkedList.LinkedListPrependAdder{})
	linkedList1 = ArrayToList([]int{4, 3, 1}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]int{4, 1}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]int{1, 4, 8}, *NewLinkedListAppendToTailAdder())
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`SumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]int{7, 1, 6}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]int{5, 9, 2}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]int{9, 1, 2}, *NewLinkedListAppendToTailAdder())
	result = SumLists(linkedList1, linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`SumLists() expect valid linkedList`)
	}
}

func TestSumListsForwardOrder(t *testing.T) {
	linkedList1 := ArrayToList([]int{1, 3, 4}, *NewLinkedListAppendToTailAdder())
	linkedList2 := ArrayToList([]int{1, 4}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList := ArrayToList([]int{1, 4, 8}, *NewLinkedListAppendToTailAdder())
	result := SumListsForwardOrder(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumListsForwardOrder() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]int{6, 1, 7}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]int{2, 9, 5}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]int{9, 1, 2}, *NewLinkedListAppendToTailAdder())
	result = SumListsForwardOrder(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumListsForwardOrder() expect valid linkedList`)
	}
}

func TestIsPalindrome(t *testing.T) {
	newlinkedList := &linkedList.Node{"D", &linkedList.Node{"a", &linkedList.Node{"b", &linkedList.Node{"a", &linkedList.Node{"l",
		&linkedList.Node{"e", &linkedList.Node{"a", &linkedList.Node{"r", &linkedList.Node{"r", &linkedList.Node{"o", &linkedList.Node{"z",
			&linkedList.Node{"a", &linkedList.Node{"l", &linkedList.Node{"a", &linkedList.Node{"z", &linkedList.Node{"o",
				&linkedList.Node{"r", &linkedList.Node{"r", &linkedList.Node{"a", &linkedList.Node{"e", &linkedList.Node{"l",
					&linkedList.Node{"a", &linkedList.Node{"b", &linkedList.Node{"a", &linkedList.Node{"d", nil}}}}}}}}}}}}}}}}}}}}}}}}}
	result := IsPalindrome(newlinkedList)
	if result != true {
		t.Fatalf(`TestIsPalindrome() expect true`)
	}
	newlinkedList = &linkedList.Node{"P", &linkedList.Node{"e", &linkedList.Node{"r", &linkedList.Node{"r", &linkedList.Node{"o", nil}}}}}
	result = IsPalindrome(newlinkedList)
	if result != false {
		t.Fatalf(`TestIsPalindrome() expect false`)
	}
}

func TestIsIntersection(t *testing.T) {
	intersectionNode := &linkedList.Node{"d", &linkedList.Node{"a", &linkedList.Node{"s", nil}}}
	linkedList1 := &linkedList.Node{"d", &linkedList.Node{"a", &linkedList.Node{"z", intersectionNode}}}
	linkedList2 := &linkedList.Node{"p", &linkedList.Node{"e", &linkedList.Node{"r", &linkedList.Node{"r", &linkedList.Node{"o", intersectionNode}}}}}
	result := IsIntersection(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, intersectionNode) {
		t.Fatalf(`TestIsIntersection() expect a valid linked list`)
	}
	linkedList1 = &linkedList.Node{"a", &linkedList.Node{"b", &linkedList.Node{"c", nil}}}
	linkedList2 = &linkedList.Node{"z", &linkedList.Node{"b", &linkedList.Node{"c", nil}}}
	result = IsIntersection(linkedList1, linkedList2)
	if result != nil {
		t.Fatalf(`TestIsIntersection() expect nil`)
	}
}

func TestLoopDetection(t *testing.T) {
	newLinkedList := &linkedList.Node{"d", &linkedList.Node{"a", &linkedList.Node{"z", nil}}}
	newLinkedList.Next.Next.Next = newLinkedList
	result := LoopDetection(newLinkedList)
	if !reflect.DeepEqual(result.Next, newLinkedList) {
		t.Fatalf(`TestLoopDetection() expect a valid linked list`)
	}
	tail := &linkedList.Node{"d", &linkedList.Node{"a", &linkedList.Node{"z", nil}}}
	newLinkedList = &linkedList.Node{"d", &linkedList.Node{"a", &linkedList.Node{"z", nil}}}
	newLinkedList.Next.Next.Next = tail
	result = LoopDetection(newLinkedList)
	if result != nil {
		t.Fatalf(`TestIsIntersection() expect nil`)
	}
}

func TestIncrementalNumberLinkedList(t *testing.T) {
	expectedLinkedList := linkedList.Node{Data: 1, Next: nil}
	expectedLinkedList.AppendToTail(2)
	expectedLinkedList.AppendToTail(3)
	expectedLinkedList.AppendToTail(4)
	expectedLinkedList.AppendToTail(5)
	expectedLinkedList.AppendToTail(6)
	expectedLinkedList.AppendToTail(7)

	response := CreateIncrementalNumberLinkedList(1, 7)
	if !reflect.DeepEqual(response, expectedLinkedList) {
		t.Fatalf(`CreateIncrementalNumberLinkedList(linkedList, 2) expect a valid linked list`)
	}
	response = CreateIncrementalNumberLinkedList(4, 2)
	if !reflect.DeepEqual(response, linkedList.Node{}) {
		t.Fatalf(`CreateIncrementalNumberLinkedList(linkedList, 2) expect a valid linked list`)
	}
	expectedLinkedList = linkedList.Node{Data: 6, Next: nil}
	response = CreateIncrementalNumberLinkedList(6, 6)
	if !reflect.DeepEqual(response, expectedLinkedList) {
		t.Fatalf(`CreateIncrementalNumberLinkedList(linkedList, 2) expect a valid linked list`)
	}
}

func TestNewLinkedListAppendToTailAdder(t *testing.T) {
	linkedListAdder := linkedList.NewLinkedListAdder(linkedList.LinkedListAppendToTailAdder{})
	if *linkedListAdder != *NewLinkedListAppendToTailAdder() {
		t.Fatalf(`NewLinkedListAppendToTailAdder() expect a LinkedListAppendToTailAdder`)
	}
}
