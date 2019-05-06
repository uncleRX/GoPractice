package main

import "fmt"

func main() {
	// for循环
	var maxCount = 10
	var a int

	// style1
	for i := 0; i < maxCount; i++ {
		fmt.Println("i", i)
	}
	// style2
	for a < maxCount {
		a++
		fmt.Println("a:", a)
	}

	// style3
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

}
