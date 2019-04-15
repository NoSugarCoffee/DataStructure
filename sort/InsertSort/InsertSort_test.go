package InsertSort

import (
	"testing"
	"reflect"
)

func TestInsertSort(t *testing.T) {
	arr:=[]int{2,4,3,5,6,7,1}
	act:=[]int{7,6,5,4,3,2,1}
	InsertSort(arr)
	if !reflect.DeepEqual(arr,act){
		t.Fail()
	}
	arr=[]int{2,4,3,5,6,7,1}
	InsertSort2(arr)
	if !reflect.DeepEqual(arr,act){
		t.Fail()
	}
}