package bst

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Comparable interface {
	CompareTo(o interface{}) int
}

type Bst struct {
	Root *Node
	//节点个数
	size int
}
type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

func New() *Bst {
	return &Bst{Root: nil, size: 0}
}
func NewNode(e interface{}) *Node {
	return &Node{Val: e, Left: nil, Right: nil}
}
func (this *Bst) Add(e interface{}) {
	//表示添加第一个节点
	if this.size == 0 {
		node := NewNode(e)
		this.Root = node
		this.size++
	} else {
		this.add(&this.Root, e)
	}

}

func (this *Bst) add(root **Node, e interface{}) {
	// 此处要修改传入的 *Node 的地址 ，要传入 *Node 的地址
	if *root == nil {
		*root = NewNode(e)
		this.size++
		return
	}
	res, err := compare(*root, NewNode(e))
	if err != nil {
		panic(err)
	} else if res > 0 {
		this.add(&(*root).Left, e)
	} else {
		this.add(&(*root).Right, e)
	}
}

func compare(root *Node, newNode *Node) (int, error) {
	if reflect.TypeOf(root.Val) != reflect.TypeOf(newNode.Val) {
		panic("root , newNode type different")
	}
	switch res := root.Val.(type) {
	case int:
		v := newNode.Val.(int)
		return res - v, nil
	case string:
		v := newNode.Val.(string)
		return strings.Compare(res, v), nil
	case Comparable:
		v := newNode.Val.(Comparable)
		return res.CompareTo(v), nil
	}
	return 0, errors.New("compare error")
}

//
func (this *Bst) String() string {
	b := new(bytes.Buffer)
	b.WriteString("bst size=" + strconv.Itoa(this.size) + "\n")
	return b.String() + qian(this.Root)
}

//先序遍历
func qian(root *Node) string {
	var b bytes.Buffer
	if root == nil {
		return "nil "
	}
	switch res := root.Val.(type) {
	case int:
		b.WriteString(strconv.Itoa(root.Val.(int)) + " ")
	case string:
		b.WriteString(res + " ")
	case Comparable:
		b.WriteString(fmt.Sprintf("%v ", res))
	}
	return b.String() + qian(root.Left) + qian(root.Right)
}
