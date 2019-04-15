package selectSort

import "fmt"

func SelectSort(arr []int) {
	for i := range arr {
		max := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		if max != i{
			arr[max],arr[i]=arr[i],arr[max]
		}
		fmt.Printf("第%d次:%v\n",i,arr[:])
	}
}
