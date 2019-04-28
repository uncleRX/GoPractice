package main

import "fmt"

func main() {

	// 声明一个Go指针
	testPoint()
	testNilPoint()
	testPointArr()
	testPP()
	testPPByParam()
}

// 指针
func testPoint() {

	a := 10
	var ip *int
	ip = &a

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

// 空指针
func testNilPoint() {
	// 当一个指针没有分配任何变量的时候,它的值为nil
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)

	if ptr != nil {
		fmt.Printf("ptr不为空\n")
	} else {
		fmt.Printf("ptr为空\n")
	}
}

// 指针数组
func testPointArr() {

	a := []int{10, 20, 30}
	var ptr [3]*int
	for i := 0; i < 3; i++ {
		// 整数地址赋值给指针数组
		ptr[i] = &a[i]
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d] = %x\n", i, ptr[i])
	}

}

// 指向指针的指针
func testPP() {
	var a int
	var ptr *int
	var pptr **int

	a = 10
	ptr = &a
	pptr = &ptr
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

func testPPByParam() {

	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap(&a, &b)
	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
