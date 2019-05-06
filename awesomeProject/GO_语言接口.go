package main

import "fmt"

// Go 语言提供了另外一种数据类型 - 接口 ,他把所有具有共性的方法定义在一起,任何其他类型只要实现了这些方法就是实现了这个接口


// 声明一个人的接口
type People interface {
	speak()
}


type Doctor struct {

}

func (Doctor) speak() {
	fmt.Println("Doctor - speak")
}


func main() {

 	var doc = new(Doctor)
 	doc.speak()
 	//var pro = new(Programer)

	var p1 People
 	p1 = doc
	p1.speak()


}