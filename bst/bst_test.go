package bst

import (
	"testing"
)

type Student struct {
	Id   int
	Name string
}

// 1:this>s   0:=   -1:this<s
func (this *Student) CompareTo(s interface{}) int {
	return this.Id - s.(*Student).Id
}

func TestBst(t *testing.T) {
	bst := New()
	bst.Add(1)
	bst.Add(2)
	bst.Add(0)
	bst.Add(5)
	t.Log(bst)

	bst2 := New()
	bst2.Add("aa")
	bst2.Add("bb")
	bst2.Add("cc")
	bst2.Add("bc")
	t.Log(bst2)

	stus := New()
	stus.Add(&Student{Id: 1, Name: "aa"})
	stus.Add(&Student{Id: 2, Name: "bb"})
	stus.Add(&Student{Id: 0, Name: "cc"})
	stus.Add(&Student{Id: 5, Name: "dd"})
	t.Log(stus)
}
