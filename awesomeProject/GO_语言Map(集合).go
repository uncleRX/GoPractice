package main

import (
	"fmt"
)

func main() {
	// map: 无需的键值对的集合,map是使用hash表来实现的

	// 声明一个map style1
	var map1 map[string]string
	var map2 = make(map[string]string)

	fmt.Println(map1,map2)

	// 插入键值对
	//map1["name"] = "小花" error: assignment to entry in nil map
	map2["name"] = "小李"
	fmt.Println(map1,map2)

	// 查看元素是否存在
	name, isOk := map2 ["name"]
	if isOk {
		fmt.Println("name存在:",name)

	}else {
		fmt.Println("name不存在")
	}

	delete(map2,"name" )

	name1, isOk1 := map2 ["name"]
	if isOk1 {
		fmt.Println("name存在:",name1)

	}else {
		fmt.Println("name不存在")
	}





	// delete


}