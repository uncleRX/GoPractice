package main

import "fmt"

func main() {

	// Go 语言中 range 关键字用于 for 循环中迭代array,切片,map,channel

	// 声明一个切片
	nums := []int{1,2,3}
	sum := 0
	for _,nums := range nums {
		sum += nums
	}
	fmt.Println(sum)


}