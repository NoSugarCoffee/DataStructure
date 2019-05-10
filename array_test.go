package array

import (
	"testing"
)

type Student struct {
	Name string
}

func (this *Student) equalse(student Student) bool {
	return this.Name == student.Name
}

func TestArray_GetSizeAndGetCapacity(t *testing.T) {
	array := NewArray(10)
	t.Log(array)
	if array.IsEmpty() != true {
		t.Error("should empty")
	}
	if ok := array.Add(0, 1); !ok {
		t.Error("add error")
	}
	t.Log(array)
	array.AddLast(3)
	t.Log(array)
	array.AddFirst(0)
	t.Log(array)
	if array.Get(0) != 0 {
		t.Log("array.Get error")
	}
	array.Add(2, 2)
	t.Log(array)
	if res := array.Remove(1); res != 1 {
		t.Log(res, "remove error")
	}
	t.Log(array)
	array.RemoveFirst()
	t.Log(array)
	array.RemoveLast()
	t.Log(array)
	array.AddLast(2)
	t.Log(array)
	array.RemoveAllElement(2)
	t.Log(array)

}
