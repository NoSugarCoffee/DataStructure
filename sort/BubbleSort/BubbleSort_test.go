package BubbleSort

import (
	"testing"
	"reflect"
)

func TestBubbleSort(t *testing.T) {
	arr:=[]int{2,4,3,5,6,7,1}
	act:=[]int{7,6,5,4,3,2,1}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr,act){
		t.Fail()
	}
	arr=[]int{2,4,3,5,6,7,1}
	BubbleSort2(arr)
	if !reflect.DeepEqual(arr,act){
		t.Fail()
	}
}
