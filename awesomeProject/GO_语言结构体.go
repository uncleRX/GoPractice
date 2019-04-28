package main

import "fmt"

// 定义结构体
type People struct {
	name string
	age  int
}

func main() {
	// 创建一个结构体
	var man1 = People{"小花", 10}
	fmt.Println("man1_name:", man1.name, "man1_age:", man1.age)

	var man2 = People{}
	man2.name = "小花2"
	man2.age = 20
	fmt.Println("man2_name:", man2.name, "man2_age:", man2.age)
	logPeople(man1)

}

// 结构体作为参数传递
func logPeople(man People) {
	fmt.Println("this_name:", man.name, "this_age:", man.age)

}

// 结构体指针
func testPointStruct() {

}
