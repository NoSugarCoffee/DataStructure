package array

import (
	"fmt"
	"reflect"
	"strings"
)

type Array struct {
	data []interface{}
}

var defaultCapacity int = 10

// 构造函数
func New(capacity int) *Array {
	if capacity < 1 {
		panic("capacity can't less than 1")
	}
	return &Array{data: make([]interface{}, 0, capacity)}
}

func NewWithDefaultCapacity() *Array {
	return New(defaultCapacity)
}

func (this *Array) GetSize() int {
	return len(this.data)
}

func (this *Array) GetCapacity() int {
	return cap(this.data)
}

func (this *Array) IsEmpty() bool {
	return this.GetSize() == 0
}
func (this *Array) Add(index int, e interface{}) bool {
	if index < 0 || index > this.GetSize() {
		panic("index error")
	}
	tempSlice := make([]interface{}, this.GetSize()-index)
	copy(tempSlice, this.data[index:])
	this.data = append(append(this.data[:index], e), tempSlice...)
	return true
}
func (this *Array) AddLast(e interface{}) {
	this.data = append(this.data, e)
}
func (this *Array) AddFirst(e interface{}) {
	this.Add(0, e)
}

func (this *Array) Remove(index int) interface{} {
	if index < 0 || index > this.GetSize()-1 {
		panic("index error")
	}
	rmItem := this.data[index]
	this.data = append(this.data[0:index], this.data[index+1:]...)
	return rmItem
}
func (this *Array) RemoveFirst() interface{} {
	return this.Remove(0)
}
func (this *Array) RemoveLast() interface{} {
	return this.Remove(this.GetSize() - 1)
}
func (this *Array) Get(index int) interface{} {
	if index < 0 || index > this.GetSize()-1 {
		panic("index error")
	}
	return this.data[index]
}
func (this *Array) GetFirst() interface{} {
	return this.Get(0)
}

func (this *Array) GetLast() interface{} {
	return this.Get(this.GetSize() - 1)
}

func (this *Array) RemoveAllElement(e interface{}) {
	for _, v := range this.data {
		this.RemoveElement(v)
	}
}
func (this *Array) RemoveElement(e interface{}) (res interface{}) {
	for i, v := range this.data {
		if reflect.DeepEqual(v, e) {
			res = this.Remove(i)
			break
		}
	}
	return
}
func (this *Array) Set(index int, e interface{}) {
	if index < 0 || index > this.GetSize()-1 {
		panic("index error")
	}
	this.data[index] = e
}

func (this *Array) String() string {
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
