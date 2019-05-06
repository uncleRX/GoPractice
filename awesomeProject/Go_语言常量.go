package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
	// 只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	const name = "小花"
	fmt.Println(name)

	// 支持多常量声明
	const name1, name2 = "name1", "name2"
	fmt.Println(name1, name2)

	// 常量可以用作枚举
	const (
		Unknow = 0
		Female = 1
		Male   = 2
	)
	fmt.Println("枚举Unnow的值", Unknow)

	// 值还可以是内置函数,必须是内置函数,否则会报错
	const (
		myName     = "小花花"
		nameLen    = len(myName)
		unsafeSize = unsafe.Sizeof(myName)
	)

	fmt.Println("name:", myName, "len:", nameLen, "unsafeSize:", unsafeSize)

	// iota 特殊的常量 ,可以被编译器修改的常量。
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Println(a, b, c)

	// 简写
	const (
		a1 = iota // 0
		b1        // 1
		c1        // 2
	)
	fmt.Println(a1, b1, c1)

	//iota 用法
	const (
		a2 = iota //0
		b2        //1
		c2        //2
		d2 = "ha" //独立值，iota += 1
		e2        //"ha"   iota += 1
		f2 = 100  //iota +=1
		g2        //100  iota +=1
		h2 = iota //7,恢复计数
		i2        //8
		_         //9
		_         //10
		x2        //11
	)

	// 每次 const 出现时，都会让 iota 初始化为0 , iota代表行索引
	fmt.Println(a2, b2, c2, d2, e2, f2, g2, h2, i2, x2)
}
