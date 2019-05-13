package queue

import (
	"DataStructure/array"
	"errors"
	"fmt"
	"strings"
)

type LoopQueue struct {
	data        *array.Array
	front, tail int
	size        int
	capacity    int
}

func NewLoopQueue(capacity int) *LoopQueue {
	return &LoopQueue{data: array.New(capacity + 1), front: 0, tail: 0, capacity: capacity}
}

func (this *LoopQueue) GetSize() int {
	return this.size
}

func (this *LoopQueue) GetCapacity() int {
	return this.capacity
}

func (this *LoopQueue) IsEmpty() bool {
	return this.front == this.tail
}

func (this *LoopQueue) Enqueue(e interface{}) error {
	if (this.tail+1)%this.data.GetCapacity() == this.front {
		return errors.New("queue full can't enqueue")
	}
	this.data.Add(this.tail, e)
	this.tail++
	this.tail = (this.tail) % this.data.GetCapacity()
	this.size++
	return nil
}

func (this *LoopQueue) String() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(fmt.Sprintf("LoopQueue size:%d ,capacity:%d head[", this.GetSize(), this.GetCapacity()))

	for i := this.front; i != this.tail; i = i % (this.data.GetCapacity()) {
		stringBuilder.WriteString(fmt.Sprint(this.data.Get(i)))
		if (i+1)%this.data.GetCapacity() != this.tail {
			stringBuilder.WriteString(",")
		}
		i++
	}
	stringBuilder.WriteString("]tail")
	return stringBuilder.String()
}

func (this *LoopQueue) Dequeue() (interface{}, error) {
	if this.front == this.tail {
		return nil, errors.New("queue empty ")
	}
	ret := this.data.Get(this.front)
	this.front++
	this.front = (this.front) % this.data.GetCapacity()
	this.size--
	return ret, nil
}
