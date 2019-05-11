package queue

import (
	"DataStructure/array"
	"fmt"
	"strings"
)

type Queue struct {
	data *array.Array
}

func New(capacity int) *Queue {
	return &Queue{data: array.New(capacity)}
}

func NewWithDefaultCapacity() *Queue {
	return &Queue{data: array.NewWithDefaultCapacity()}
}

func (this *Queue) GetSize() int {
	return this.data.GetSize()
}

func (this *Queue) GetCapacity() int {
	return this.data.GetCapacity()
}

func (this *Queue) IsEmpty() bool {
	return this.data.IsEmpty()
}

func (this *Queue) Enqueue(e interface{}) {
	this.data.AddLast(e)
}

func (this *Queue) Dequeue() interface{} {
	return this.data.RemoveFirst()
}
func (this *Queue) GetFront() interface{} {
	return this.data.GetFirst()
}
func (this *Queue) String() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(fmt.Sprintf("Queue size:%d ,capacity:%d head[", this.GetSize(), this.GetCapacity()))

	for i := 0; i < this.GetSize(); i++ {
		stringBuilder.WriteString(fmt.Sprint(this.data.Get(i)))
		if i != this.GetSize()-1 {
			stringBuilder.WriteString(",")
		}
	}
	stringBuilder.WriteString("]tail")
	return stringBuilder.String()
}
