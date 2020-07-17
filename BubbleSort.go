package main

import "fmt"

func main()  {
	var ary = []int{12,89,29,28,37,11,76,24,56,99}
	arylen := len(ary)
	fmt.Println("未排序前的顺序为：",ary)
	for i := 0; i < arylen; i++ {
		for j := i; j < arylen; j++ {
			if ary[i] > ary[j] {
				temp := ary[i]
				ary[i] = ary[j]
				ary[j] = temp
			}
		}
	}
	fmt.Println("排序后的顺序为：",ary)
}

