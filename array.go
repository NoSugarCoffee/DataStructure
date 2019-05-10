package array

import (
	"fmt"
	"reflect"
	"strings"
)

type array struct {
	data []interface{}
}

var defaultCapacity int = 10

// 构造函数
func NewArray(capacity int) *array {
	if capacity < 1 {
		panic("capacity can't less than 1")
	}
	return &array{data: make([]interface{}, 0, capacity)}
}

func NewArrayWithDefaultCapacity() *array {
	return NewArray(defaultCapacity)
}

func (this *array) GetSize() int {
	return len(this.data)
}

func (this *array) GetCapacity() int {
	return cap(this.data)
}

func (this *array) IsEmpty() bool {
	return this.GetSize() == 0
}
func (this *array) Add(index int, e interface{}) bool {
	if index < 0 || index > this.GetSize() {
		panic("index error")
	}
	tempSlice := make([]interface{}, this.GetSize()-index)
	copy(tempSlice, this.data[index:])
	this.data = append(append(this.data[:index], e), tempSlice...)
	return true
}
func (this *array) AddLast(e interface{}) {
	this.data = append(this.data, e)
}
func (this *array) AddFirst(e interface{}) {
	this.Add(0, e)
}

func (this *array) Remove(index int) interface{} {
	if index < 0 || index > this.GetSize()-1 {
		panic("index error")
	}
	rmItem := this.data[index]
	this.data = append(this.data[0:index], this.data[index+1:]...)
	return rmItem
}
func (this *array) RemoveFirst() interface{} {
	return this.Remove(0)
}
func (this *array) RemoveLast() interface{} {
	return this.Remove(this.GetSize() - 1)
}
func (this *array) Get(index int) interface{} {
	if index < 0 || index > this.GetSize()-1 {
		panic("index error")
	}
	return this.data[index]
}

func (this *array) RemoveAllElement(e interface{}) {
	for _, v := range this.data {
		this.RemoveElement(v)
	}
}
func (this *array) RemoveElement(e interface{}) (res interface{}) {
	for i, v := range this.data {
		if reflect.DeepEqual(v, e) {
			res = this.Remove(i)
			break
		}
	}
	return
}
func (this *array) Set(index int, e interface{}) {
	if index < 0 || index > this.GetSize()-1 {
		panic("index error")
	}
	this.data[index] = e
}

func (this *array) String() string {
	var stringBuilder strings.Builder
	stringBuilder.WriteString(fmt.Sprintf("Array size:%d ,capacity:%d [", this.GetSize(), this.GetCapacity()))
	for i, item := range this.data {
		stringBuilder.WriteString(fmt.Sprint(item))
		if i != this.GetSize()-1 {
			stringBuilder.WriteString(",")
		}
	}
	stringBuilder.WriteString("]")
	return stringBuilder.String()
}
