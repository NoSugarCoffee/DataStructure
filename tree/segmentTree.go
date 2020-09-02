package tree

import (
	"fmt"
)

type Node struct {
	Left,Right int
	Value int
	Index int
}
type SegmentTree struct {
	// 结点信息
	nodes []*Node
	// 原始数据
	arr [] int
}

func _init(arr []int) SegmentTree {
	return SegmentTree{nodes: make([]*Node, len(arr) * 4), arr:arr[:]}
}

func LeftChildIndex(treeIndex int) int {
	return treeIndex * 2 + 1
}


func RightChildIndex(treeIndex int) int {
	return treeIndex * 2 + 2
}

func (this *SegmentTree) leftChild(treeIndex int) *Node {
	return this.nodes[LeftChildIndex(treeIndex)]
}


func (this *SegmentTree) rightChild(treeIndex int) *Node {
	return this.nodes[RightChildIndex(treeIndex)]
}

// start，end: arr 的区间
// treeIndex: 线段树在数组中的位置
func(this *SegmentTree) build(start, end, treeIndex int) {
	// 叶子结点
	if start == end {
		this.nodes[treeIndex] = &Node{Left:start, Right:end, Value:this.arr[start], Index:treeIndex}
		return
	}
	mid := (start + end) / 2
	this.build(start, mid, LeftChildIndex(treeIndex))
	this.build(mid + 1, end, RightChildIndex(treeIndex))
	this.nodes[treeIndex] = &Node{Left: start, Right: end,
	Value:this.leftChild(treeIndex).Value + this.rightChild(treeIndex).Value, Index:treeIndex}
}

func (this * SegmentTree) updateSingle(index int, value int) {
	if this.nodes[0].Left > index || this.nodes[0].Right < index {
		panic(fmt.Sprintf("index out of range %d", index))
	}
	this.update(index, value, this.nodes[0])
}


func (this * SegmentTree) update(index int, value int, root *Node) {
	if root.Left == root.Right && root.Left == index {
		root.Value = value
		return
	}
	mid := (root.Left + root.Right) / 2
	if index <= mid {
		this.update(index, value, this.leftChild(root.Index))
	} else {
		this.update(index, value, this.rightChild((*root).Index))
	}
	root.Value = this.leftChild(root.Index).Value + this.rightChild(root.Index).Value
}

func (this * SegmentTree) queryRecursion(l, r int, root Node) int {
	fmt.Printf("root.index, root.value, %d,%d l:%d , r:%d\n", root.Index, root.Value ,l, r)

	var (
		left, right int
	)
	//         |-----| root
	// |------|          |------|
	if r < root.Left || l > root.Right {
		return 0
	}
	//  |---| root
	// |------|
	// 叶子结点或者是改结点是查询范围的一部分
	if l <= root.Left && root.Right <= r {
		return root.Value
	}
	left = this.queryRecursion(l, r, *this.leftChild(root.Index))
	right = this.queryRecursion(l, r, *this.rightChild(root.Index))
	return left + right
}

func (this * SegmentTree) Query (l, r int) int{
	root := *this.nodes[0]
	// 要在指定范围内查询
	if !(l >= root.Left && l <= root.Right && r >= root.Left && r <= root.Right) {
		fmt.Printf("超出查询范围[%d:%d]", root.Left,root.Right)
		return -1
	}
	return this.queryRecursion(l,r, *this.nodes[0])
}

func BuildSegmentTree (arr []int) SegmentTree {
	// 使用满二叉树的存储方式，元数据只可能会存储在最后两层
	// 用线段树存储 n 个数据，
	// 1.都存储在最后一层，满足 2^x = n，x=(0,1,2,3...) ，则需要 n-1 + n = 2n-1
	// 2.存储在两层，找到 2^x < n < 2^(x+1)
	// 2^(x+2)-1
	// 则需要多 2(2n-1) +1 = 4n-1
	seg := _init(arr)
	seg.build( 0, len(arr) - 1, 0)
	return seg
}
