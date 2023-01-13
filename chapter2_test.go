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
	newLinkedList := ArrayToList([]interface{}{3, 5, 8, 5, 10, 2, 1}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList := ArrayToList([]interface{}{1, 2, 3, 5, 8, 5, 10}, *NewLinkedListAppendToTailAdder())
	result := Partition(*newLinkedList, 5)
	if !reflect.DeepEqual(result, *expectedLinkedList) {
		t.Fatalf(`Partition() expect valid linkedList`)
	}
}

func TestSumLists(t *testing.T) {
	linkedList1 := ArrayToList([]interface{}{4, 3, 4}, *NewLinkedListAppendToTailAdder())
	linkedList2 := ArrayToList([]interface{}{1, 4}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList := ArrayToList([]interface{}{5, 7, 4}, *NewLinkedListAppendToTailAdder())
	linkedListAdder := linkedList.NewLinkedListAdder(linkedList.LinkedListAppendToTailAdder{})
	result := SumLists(*linkedList1, *linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]interface{}{2, 3}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]interface{}{1, 4, 8}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]interface{}{3, 7, 8}, *NewLinkedListAppendToTailAdder())
	result = SumLists(*linkedList1, *linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]interface{}{7, 1, 6}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]interface{}{5, 9, 2}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]interface{}{2, 1, 9}, *NewLinkedListAppendToTailAdder())
	result = SumLists(*linkedList1, *linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]interface{}{9, 9, 9}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]interface{}{9, 9, 9}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]interface{}{8, 9, 9, 1}, *NewLinkedListAppendToTailAdder())
	result = SumLists(*linkedList1, *linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumLists() expect valid linkedList`)
	}

	linkedListAdder = linkedList.NewLinkedListAdder(linkedList.LinkedListPrependAdder{})
	linkedList1 = ArrayToList([]interface{}{4, 3, 1}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]interface{}{4, 1}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]interface{}{1, 4, 8}, *NewLinkedListAppendToTailAdder())
	result = SumLists(*linkedList1, *linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`SumLists() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]interface{}{7, 1, 6}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]interface{}{5, 9, 2}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]interface{}{9, 1, 2}, *NewLinkedListAppendToTailAdder())
	result = SumLists(*linkedList1, *linkedList2, *linkedListAdder)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`SumLists() expect valid linkedList`)
	}
}

func TestSumListsForwardOrder(t *testing.T) {
	linkedList1 := ArrayToList([]interface{}{1, 3, 4}, *NewLinkedListAppendToTailAdder())
	linkedList2 := ArrayToList([]interface{}{1, 4}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList := ArrayToList([]interface{}{1, 4, 8}, *NewLinkedListAppendToTailAdder())
	result := SumListsForwardOrder(*linkedList1, *linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumListsForwardOrder() expect valid linkedList`)
	}
	linkedList1 = ArrayToList([]interface{}{6, 1, 7}, *NewLinkedListAppendToTailAdder())
	linkedList2 = ArrayToList([]interface{}{2, 9, 5}, *NewLinkedListAppendToTailAdder())
	expectedLinkedList = ArrayToList([]interface{}{9, 1, 2}, *NewLinkedListAppendToTailAdder())
	result = SumListsForwardOrder(*linkedList1, *linkedList2)
	if !reflect.DeepEqual(result, expectedLinkedList) {
		t.Fatalf(`TestSumListsForwardOrder() expect valid linkedList`)
	}
}

func TestIsPalindrome(t *testing.T) {
	newlinkedList := ArrayToList([]interface{}{
		"D", "a", "b", "a", "l", "e", "a", "r", "r", "o", "z", "a", "l", "a", "z", "o", "r", "r", "a", "e", "l", "A", "b", "a", "d",
	}, *NewLinkedListAppendToTailAdder())
	result := IsPalindrome(newlinkedList)
	if result != true {
		t.Fatalf(`TestIsPalindrome() expect true`)
	}
	newlinkedList = ArrayToList([]interface{}{"p", "e", "r", "r", "o"}, *NewLinkedListAppendToTailAdder())
	result = IsPalindrome(newlinkedList)
	if result != false {
		t.Fatalf(`TestIsPalindrome() expect false`)
	}
}

func TestIsIntersection(t *testing.T) {
	intersectionNode := &linkedList.Node{Data: "d", Next: &linkedList.Node{Data: "a", Next: &linkedList.Node{Data: "s", Next: nil}}}
	linkedList1 := &linkedList.Node{Data: "d", Next: &linkedList.Node{Data: "a", Next: &linkedList.Node{Data: "z", Next: intersectionNode}}}
	linkedList2 := &linkedList.Node{Data: "p", Next: &linkedList.Node{
		Data: "e", Next: &linkedList.Node{Data: "r", Next: &linkedList.Node{
			Data: "r", Next: &linkedList.Node{Data: "o", Next: intersectionNode}}}}}
	result := IsIntersection(linkedList1, linkedList2)
	if !reflect.DeepEqual(result, intersectionNode) {
		t.Fatalf(`TestIsIntersection() expect a valid linked list`)
	}
	linkedList1 = &linkedList.Node{Data: "a", Next: &linkedList.Node{Data: "b", Next: &linkedList.Node{Data: "c", Next: nil}}}
	linkedList2 = &linkedList.Node{Data: "z", Next: &linkedList.Node{Data: "b", Next: &linkedList.Node{Data: "c", Next: nil}}}
	result = IsIntersection(linkedList1, linkedList2)
	if result != nil {
		t.Fatalf(`TestIsIntersection() expect nil`)
	}
}

func TestLoopDetection(t *testing.T) {
	newLinkedList := ArrayToList([]interface{}{"d", "a", "z"}, *NewLinkedListAppendToTailAdder())
	newLinkedList.Next.Next.Next = newLinkedList
	result := LoopDetection(newLinkedList)
	if !reflect.DeepEqual(result.Next, newLinkedList) {
		t.Fatalf(`TestLoopDetection() expect a valid linked list`)
	}
	tail := ArrayToList([]interface{}{"d", "a", "z"}, *NewLinkedListAppendToTailAdder())
	newLinkedList = ArrayToList([]interface{}{"d", "a", "z"}, *NewLinkedListAppendToTailAdder())
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
