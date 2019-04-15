package BubbleSort

import "fmt"

// 冒泡算法特点:
// 1.每次都有一个数字到达最终位置，若从大到小，则每次都会有最小数字在右侧
// 2.复杂度 O(n*n)
func BubbleSort(arr []int) {
	//n个数 只要排序 n-1次
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}

		}
		fmt.Printf("第%d次:%v\n",i,arr[:])
	}
}

// 冒泡排序改进,减少已经有序的次数
func BubbleSort2(arr []int){
	for i:=1;i<len(arr);i++{
		flag:=true
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] < arr[j+1] {
				flag=false
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
		if flag{
			break
		}
		fmt.Printf("第%d次:%v\n",i,arr[:])
	}
}
