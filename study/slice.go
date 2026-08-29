package main

import "fmt"

// slice 函数演示了Go语言中切片的基本操作和append函数的使用
func slice() {
	// 创建一个包含1, 2, 3的整型切片slice1
	slice1 := []int{1, 2, 3}
	/*
		创建一个添加4, 5到slice1数组的整型切片slice2
		和slice2 := append(slice1, 4, 5)等价
	*/
	var slice2 []int = append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 3, 6)
	fmt.Println(slice3)
	fmt.Println(slice3[:cap(slice3)])

	slice4 := append(slice3, 4, 5)
	slice4[0] = 10
	fmt.Println(slice3)
	fmt.Println(slice4[:cap(slice4)])
	// 以上表明，当slice3和slice4指向同一个底层数组时，修改slice4会影响到slice3

	slice5 := make([]int, 3, 6)
	slice6 := append(slice5[0:3:3], 4, 5)
	fmt.Println(slice5)
	fmt.Println(slice6[:cap(slice6)])

	slice5[0] = 10
	fmt.Println(slice5)
	fmt.Println(slice6[:cap(slice6)])
	// 以上表明，当slice6[0:3:3]规定好在前三位后扩容时，底层数组会重新分配，slice6和slice5指向的底层数组不同
}
