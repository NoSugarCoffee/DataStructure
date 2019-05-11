package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewWithDefaultCapacity()
	t.Log(stack)
	stack.Push(1)
	t.Log(stack)
	if res := stack.Peek(); res != 1 {
		t.Error("peek error")
	}
	stack.Push(2)
	t.Log(stack)
	if res := stack.Peek(); res != 2 {
		t.Error("peek error")
	}
}
