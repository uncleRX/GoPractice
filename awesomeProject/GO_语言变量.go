package main

import "fmt"

// 全局变量

var (
	a int
	b bool
)

func main() {
	// 初始化一个变量
	var name string = "abc"
	var index int
	fmt.Println("name:", name, "index:", index)

	// 根据值自行判断变量的类型
	var v_name = "小花"
	fmt.Println("v_name:", v_name)

	// := 需要声明 新变量,可以省略 var 关键字
	v_name1 := "小哦"
	fmt.Println("v_name1:", v_name1)

	// 声明多变量
	var vname1, vname2 = "小强", "小勇"
	fmt.Println(vname1, vname2)

	// := 声明多个变量

	vname3, vname4 := "A", "B"
	fmt.Println(vname3, vname4)

}
