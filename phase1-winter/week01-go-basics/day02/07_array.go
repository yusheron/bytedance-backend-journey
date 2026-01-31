package main

import "fmt"

func main() {
	fmt.Println("=== 数组（定长）===\n")

	// 声明并初始化
	var arr1 [5]int // 未初始化，默认零值[0 0 0 0 0]
	fmt.Printf("arr1: %v\n", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr2: %v\n", arr2)

	// 部分初始化
	arr3 := [5]int{1, 2} // 剩余元素为零值
	fmt.Printf("arr3: %v\n", arr3)

	// 自动推导长度
	arr4 := [...]int{10, 20, 30}
	fmt.Printf("arr4: %v (长度: %d)\n\n", arr4, len(arr4))

	// 访问元素
	fmt.Println("--- 访问元素 ---")
	fmt.Printf("arr2[0] = %d\n", arr2[0])
	fmt.Printf("arr2[4] = %d\n", arr2[4])
	// fmt.Println(arr2[5]) // 越界，panic

	// 修改元素
	arr2[0] = 100
	fmt.Printf("修改后: %v\n\n", arr2)

	// 数组长度
	fmt.Printf("数组长度: %d\n\n", len(arr2))

	// 遍历数组
	fmt.Println("--- 遍历数组 ---")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %d\n", i, arr2[i])
	}
	fmt.Println()

	// 用range遍历
	for index, value := range arr2 {
		fmt.Printf("index=%d, value=%d\n", index, value)
	}
	fmt.Println()

	// 数组是值类型（重要！）
	fmt.Println("--- 数组是值类型 ---")
	a := [3]int{1, 2, 3}
	b := a // 复制整个数组
	b[0] = 100
	fmt.Printf("a: %v\n", a) // [1 2 3]
	fmt.Printf("b: %v\n", b) // [100 2 3]
	fmt.Println("数组传递会复制整个数组，性能较差")
}
