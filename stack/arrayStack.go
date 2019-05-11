package stack

import (
	"DataStructure/array"
	"fmt"
	"strings"
)

type Stack struct {
	data *array.Array
}

func New(capacity int) *Stack {
	return &Stack{data: array.New(capacity)}
}
func NewWithDefaultCapacity() *Stack {
	return &Stack{data: array.NewWithDefaultCapacity()}
}
func (this *Stack) GetSize() int {
	return this.data.GetSize()
}
func (this *Stack) GetCapacity() int {
	return this.data.GetCapacity()
}

func (this *Stack) IsEmpty() bool {
	return this.IsEmpty()
}

func (this *Stack) Push(e interface{}) {
	this.data.AddLast(e)
}

func (this *Stack) Pop() interface{} {
	return this.data.RemoveLast()
}

func (this *Stack) Peek() interface{} {
	return this.data.Get(this.GetSize() - 1)
}
func (this *Stack) Get(index int) interface{} {
	return this.data.Get(index)
}

func (this *Stack) String() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(fmt.Sprintf("Stack size:%d ,capacity:%d bottom[", this.GetSize(), this.GetCapacity()))

	for i := 0; i < this.GetSize(); i++ {
		stringBuilder.WriteString(fmt.Sprint(this.Get(i)))
		if i != this.GetSize()-1 {
			stringBuilder.WriteString(",")
		}
	}
	stringBuilder.WriteString("]top")
	return stringBuilder.String()
}
