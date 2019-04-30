package main

import (
	"fmt"
)

func main() {
	// Go 语言切片是对数组的抽象。可以看做 动态数组的长度是动态的,切片是引用类型
	// 定义切片style1
	var slice1 []int // 定义了一个空切片
	fmt.Println(slice1)

	// 定义切片style2: 使用make()函数创建
	var slice2 []int = make([]int, 10)
	slice2 = append(slice2, 1)
	fmt.Println(slice2)

	// tip: []表示是切片类型

	// 切片初始化
	slice3 := []int{1, 2, 3}
	fmt.Println(slice3)

	// append() 函数
	slice3 = append(slice3, 4)

	// copy()
	// srcSlice 为数据来源切片。
	// destSlice 为复制的目标。目标切片必须分配过空间且足够承载复制的元素个数。来源和目标的类型一致，copy 的返回值表示实际发生复制的元素个数。
	var slice4 = []int{110}
	fmt.Println(slice4) // [110]
	copy(slice3,slice4)
	copy(slice4,slice3)

	fmt.Println(slice3) // [110 2 3 4]
	fmt.Println(slice4) // [110]

	// 切片的截取
	var slice5 = []int {1,2,3,4,5,6}
	fmt.Println(slice5[0:6]) // [1 2 3 4 5 6]
	fmt.Println(slice5[1:5]) // [2 3 4 5]
	fmt.Println(slice5[:5]) // [1 2 3 4 5]
	fmt.Println(slice5[0:]) // [1 2 3 4 5 6]
	fmt.Println(slice5[:]) //[1 2 3 4 5 6]




}
