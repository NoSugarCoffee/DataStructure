package trees

import (
	"fmt"
	"math"
)

type Node struct {
	Left, Right int
	Value       int
	Index       int
}
type SegmentTree struct {
	// 结点信息
	nodes []*Node
}

func _init(arr []int) *SegmentTree {
	return &SegmentTree{nodes: make([]*Node, len(arr)*4)}
}

func LeftChildIndex(treeIndex int) int {
	return treeIndex*2 + 1
}

func RightChildIndex(treeIndex int) int {
	return treeIndex*2 + 2
}

func (this *SegmentTree) leftChild(treeIndex int) *Node {
	return this.nodes[LeftChildIndex(treeIndex)]
}

func (this *SegmentTree) rightChild(treeIndex int) *Node {
	return this.nodes[RightChildIndex(treeIndex)]
}

// start，end: arr 的区间
// treeIndex: 线段树在数组中的位置
func (this *SegmentTree) build(arr []int, start, end, treeIndex int) {
	// 叶子结点存值
	if start == end {
		this.nodes[treeIndex] = &Node{Left: start, Right: end, Value: arr[start], Index: treeIndex}
		return
	}
	mid := (start + end) / 2
	this.build(arr, start, mid, LeftChildIndex(treeIndex))
	this.build(arr, mid+1, end, RightChildIndex(treeIndex))
	// 根据业务，此处是 sum
	this.nodes[treeIndex] = &Node{Left: start, Right: end,
		Value: this.leftChild(treeIndex).Value + this.rightChild(treeIndex).Value, Index: treeIndex}
}

// 可以直接调用 call  UpdateRange()
// arr[index] = value
func (this *SegmentTree) UpdateSingle(index int, value int) {
	if this.nodes[0].Left > index || this.nodes[0].Right < index {
		panic(fmt.Sprintf("index out of range %d", index))
	}
	this.updateRecursion(index, value, this.nodes[0])


}

func (this *SegmentTree) updateRecursion(index int, value int, root *Node) {
	if  root.Left == index {
		root.Value = value
		return
	}
	mid := (root.Left + root.Right) / 2
	if index <= mid {
		this.updateRecursion(index, value, this.leftChild(root.Index))
	} else {
		this.updateRecursion(index, value, this.rightChild((*root).Index))
	}
	root.Value = this.leftChild(root.Index).Value + this.rightChild(root.Index).Value
}

func (this *SegmentTree) UpdateRange(l , r int, value int) {
	root := *this.nodes[0]
	if !(l >= root.Left && l <= root.Right && r >= root.Left && r <= root.Right) {
		fmt.Printf("超出查询范围[%d:%d]", root.Left, root.Right)
	}
	this.updateRangeRecursion(l, r, this.nodes[0] ,value)
}

func (this *SegmentTree) updateRangeRecursion(l, r int, root *Node, value int) {
	//         |-----| root
	// |------|          |------|
	// 无交集直接返回
	if r < root.Left || l > root.Right {
		return
	}
	if  root.Left == root.Right  {
		root.Value = value
		return
	}
	this.updateRangeRecursion(l, r, this.leftChild(root.Index), value)
	this.updateRangeRecursion(l, r, this.rightChild(root.Index), value)
	root.Value = this.leftChild(root.Index).Value + this.rightChild(root.Index).Value
}


func (this *SegmentTree) queryRecursion(l, r int, root Node) int {
	//         |-----| root
	// |------|          |------|
	// 无交集直接返回
	if r < root.Left || l > root.Right {
		return 0
	}
	//  |---| root
	// |------|
	//区间是查询范围的一部分,无需全部遍历到叶子结点
	if l <= root.Left && root.Right <= r {
		return root.Value
	}
	sumOfLeft := this.queryRecursion(l, r, *this.leftChild(root.Index))
	sumOfRight := this.queryRecursion(l, r, *this.rightChild(root.Index))
	return sumOfLeft + sumOfRight
}

// 区间查询
func (this *SegmentTree) Query(l, r int) int {
	root := *this.nodes[0]
	// 要在指定范围内查询
	if !(l >= root.Left && l <= root.Right && r >= root.Left && r <= root.Right) {
		fmt.Printf("超出查询范围[%d:%d]", root.Left, root.Right)
		return -1
	}
	return this.queryRecursion(l, r, *this.nodes[0])
}

func BuildSegmentTree(arr []int) *SegmentTree {
	// 使用满二叉树的存储方式，元数据只可能会存储在最后两层
	// 用线段树存储 n 个数据，
	// 1.都存储在最后一层，满足 2^x = n，x=(0,1,2,3...) ，则需要 n-1 + n = 2n-1
	// 2.存储在两层，找到 2^x < n < 2^(x+1)
	seg := _init(arr)
	seg.build(arr, 0, len(arr)-1, 0)
	return seg
}

// 计算树的深度
func (this *SegmentTree) Depth() int{
	var depth int
	for depth = 1; math.Pow(2, float64(depth)) < float64(len(this.nodes)); {
		depth ++
	}
	depth --
	return depth
}

// 两个空格 == 1个2位数 == printCounter
func (this *SegmentTree) print() {
	index := 0
	for i := 0; i < this.Depth(); i++ {
		var prevPrintCounter, printCounter int
		for j := 0; j < int(math.Pow(2, float64(i))); j++ {
			var k int
			if index >= len(this.nodes) {
				break
			}
			if this.nodes[index] == nil {
				index++
				continue
			}
			// 根据线段长度打印空格 l + (r-l+0.5) / 2
			for k = 0; k < this.nodes[index].Right - this.nodes[index].Left + this.nodes[index].Left*2 - prevPrintCounter; k++ {
				fmt.Print("  ")
				printCounter++
			}
			if(this.nodes[index].Value < 10) {
				fmt.Printf("0%d", this.nodes[index].Value)
			}else {
				fmt.Printf("%d", this.nodes[index].Value)
			}
			printCounter++
			prevPrintCounter = printCounter
			index++
		}
		fmt.Print("\n")
	}
}
