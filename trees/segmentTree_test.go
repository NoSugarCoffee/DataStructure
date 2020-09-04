package trees

import (
	"fmt"
	"testing"
)

func TestBuildSegmentTree(t *testing.T) {
	// build
	arr := []int{1, 3, 5, 7, 9, 11}
	tree := BuildSegmentTree(arr)
	tree.print()

	// 区间查询
	fmt.Printf("tree.Query:,sum of [2,5] is %d\n",tree.Query(2,5))

	// 单点更新
	tree.UpdateSingle(4, 6)
	fmt.Println("update arr[4] = 6")
	tree.print()

	fmt.Println("updateRange[1,2] = 0")
	tree.UpdateRange(1,2 ,0)
	tree.print()

}