package queue

import "testing"

func TestLoopQueue(t *testing.T) {
	queue := NewLoopQueue(10)
	t.Log(queue)
	for i := 0; i < 11; i++ {
		if err := queue.Enqueue(i); err != nil {
			t.Error(err)
		}
	}
	t.Log(queue)

	for i := 0; i < 11; i++ {
		if _, err := queue.Dequeue(); err != nil {
			t.Log(err)
		}
		t.Log(queue)
	}
	t.Log(queue)
}
