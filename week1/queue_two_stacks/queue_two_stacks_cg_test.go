package queuetwostacks

import "testing"

func TestCQueue(t *testing.T) {
	// ["NewQueue", "push", "push", "peek", "pop", "empty"]
	q := NewCQueue()
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 {
		t.Error("expected 1")
	}
	if q.Pop() != 1 {
		t.Error("expected 1")
	}
	if q.Empty() {
		t.Error("expected false")
	}
}
