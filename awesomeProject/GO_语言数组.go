package main

import "fmt"

func main() {

	// Go 语言提供了数组类型的数据结构。

	// 包含同一类型且长度固定的数组,可以通过下标修改

	// 初始化数组 style1
	var floatList = [5]float32{1.00, 2.00, 3.00}
	fmt.Println(floatList)

	// style 2
	// 其元素值依次为：0，0，1，2，3。在初始化时指定了2，3，4索引中对应的值：1，2，3
	var intList = [5]int{2: 1, 3: 2, 4: 3}
	fmt.Println("intList", intList)

	// style3
	// 长度为5的数组，起元素值依次为：0，0，1，0，3。由于指定了最大索引4对应的值3，根据初始化的元素个数确定其长度为5赋值与使用
	var fiveIndexList = [...]int{2: 1, 4: 3}
	fmt.Println("fiveIndexList", fiveIndexList)

	// 注意: 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

	var moreFloatList = [...]float32{1.00, 2.00, 3.00, 1.00, 2.00, 3.00, 1.00, 2.00, 3.00}
	fmt.Println(moreFloatList)

	// 访问数组: 通过下标访问数组
	moreFloatList[0] = 1.1
	fmt.Println("修改第一个元素的值:", moreFloatList)
	firstNum := moreFloatList[0]
	fmt.Println(firstNum)
	firstNum = 2.3
	fmt.Println("引用变量修改,此数组的值:", moreFloatList)

	// 数组的迭代输出每个元素的值
	for i, x := range moreFloatList {
		fmt.Print("%v %f ", i, x)
		//fmt.Println("下标%f 值:%f",i , x)
	}

	// 迭代style2i
	for i := 0; i < len(moreFloatList); i++ {
		fmt.Printf("下标:%d 值: %f \n", i, moreFloatList[i])
	}

}
