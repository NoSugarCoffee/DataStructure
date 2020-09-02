package tree

import (
	"fmt"
	"testing"
)

func TestBuildSegmentTree(t *testing.T) {
	arr := []int {1, 3, 5, 7, 9, 11}
	tree := BuildSegmentTree(arr)
	for _,node :=range tree.nodes {
		if node != nil {
			fmt.Printf("Index %d, Sum %d for range [%d,%d]\n", (*node).Index, (*node).Value, (*node).Left, (*node).Right)
		}
	}

	tree.updateSingle(4, 6)
	fmt.Print("update\n")
	for _,node :=range tree.nodes {
		if node != nil {
			fmt.Printf("Index %d, Sum %d for range [%d,%d]\n", (*node).Index, (*node).Value, (*node).Left, (*node).Right)
		}
	}
	fmt.Print(tree.query(2,5))
}


func Test
