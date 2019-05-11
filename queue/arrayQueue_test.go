package queue

import "testing"

func TestStack(t *testing.T) {
	queue := NewWithDefaultCapacity()
	t.Log(queue)
	for i := 0; i < 11; i++ {
		queue.Enqueue(i)
	}
	t.Log(queue)
	queue.Dequeue()
	t.Log(queue)
}
