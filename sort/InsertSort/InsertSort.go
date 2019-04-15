package InsertSort

import "fmt"


//插入+交换
//排 n-1 次
func InsertSort(arr []int) {
	for i:=0;i<len(arr)-2;i++{
		// 保证找位置时数据不越界 j-1>=0
		for j:=i+1;j>0 && arr[j]>arr[j-1];j--{
			arr[j],arr[j-1]=arr[j-1],arr[j]
		}
		fmt.Printf("第%d次:%v\n",i+1,arr[:])
	}
}

//插入+选择
func InsertSort2(arr []int){
	for i:=0;i<len(arr)-2;i++{
		j:=i+1
		temp:=arr[j]
		for ;j>0 && temp>arr[j-1];j--{
			arr[j]=arr[j-1]
		}
		arr[j]=temp
		fmt.Printf("第%d次:%v\n",i+1,arr[:])
	}
}
