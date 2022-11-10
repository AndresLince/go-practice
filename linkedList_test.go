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
