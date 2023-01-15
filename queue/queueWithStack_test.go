package queue

import (
	"testing"
)

func TestIsEmptyQueue(t *testing.T) {
	queue := QueueWithStack{}
	if queue.IsEmpty() != true {
		t.Fatalf(`IsEmpty() expect true`)
	}
	queue.Add(1)
	if queue.IsEmpty() != false {
		t.Fatalf(`IsEmpty() expect false`)
	}
}

func TestPeakQueue(t *testing.T) {
	queue := QueueWithStack{}
	if queue.Peak() != nil {
		t.Fatalf(`Peak() expect nil`)
	}
	queue.Add(1)
	if queue.Peak() != 1 {
		t.Fatalf(`Peak() expect 1`)
	}
	queue.Add(2)
	if queue.Peak() != 1 {
		t.Fatalf(`Peak() expect 1`)
	}
	queue.Add(3)
	if queue.Peak() != 1 {
		t.Fatalf(`Peak() expect 1`)
	}
	queue.Remove()
	if queue.Peak() != 2 {
		t.Fatalf(`Peak() expect 2`)
	}
}

func TestPopQueue(t *testing.T) {
	queue := QueueWithStack{}
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	if queue.Remove() != 1 {
		t.Fatalf(`Remove() expect 1`)
	}
	if queue.Remove() != 2 {
		t.Fatalf(`Remove() expect 2`)
	}
	if queue.Remove() != 3 {
		t.Fatalf(`Remove() expect 3`)
	}
	if queue.Remove() != nil {
		t.Fatalf(`Remove() expect nil`)
	}
}

func TestPushQueue(t *testing.T) {
	queue := QueueWithStack{}
	queue.Add(1)
	if queue.Peak() != 1 {
		t.Fatalf(`Add() expect 1`)
	}
	queue.Add(2)
	if queue.Peak() != 1 {
		t.Fatalf(`Add() expect 1`)
	}
	queue.Add(3)
	if queue.Peak() != 1 {
		t.Fatalf(`Add() expect 1`)
	}
}
