package main

import "fmt"

func main() {

	// 对比最大值
	fmt.Println(max(1, 2))

	// 多返回值函数
	fmt.Println(swap("before", "behind"))

	// 函数传递类型
	// 1. 值传递
	// 2. 引用传递

	// GO语言中闭包
	// result 为函数类型:  func() int
	resultFunc := getAddFunc()
	// 执行
	resultValue := resultFunc()
	fmt.Println(resultValue)

}

// 普通函数(单个返回值)
// num1: 形参1 int可以省略
// num2: 形参2
// 返回值为 int
func max(num1 int, num2 int) int {
	result := num2
	if num1 > num2 {
		result = num1
	}
	return result
}

// 交换函数
// str1: 参1
// str1: 参2
// 多返回值函数
func swap(str1, str2 string) (string, string) {
	return str2, str1
}

// 返回值为函数的函数(闭包)
// func () int: 返回值为函数
func getAddFunc() func() int {
	targetCount := 0
	return func() int {
		targetCount += 1
		return targetCount
	}
}
