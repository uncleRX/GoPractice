package main

import "fmt"

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
	slice3 = append(slice3, 4)
	fmt.Println(slice3)

}
