package main

import "fmt"

func main() {

	// Go 语言中 range 关键字用于 for 循环中迭代array,切片,map,channel

	// 声明一个切片
	nums := []int{1,2,3}
	sum := 0
	// 如果不需要下标可以用 _ 代替
	for _,nums := range nums {
		sum += nums
	}
	fmt.Println(sum)

	// range可以应用于map键值对上
	maps := map[string]string{"a":"apple","b":"banana"}
	for key,value := range maps {
		fmt.Println(key,value)
	}

	// 字符串 i: 索引, c: Unicode的值
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}