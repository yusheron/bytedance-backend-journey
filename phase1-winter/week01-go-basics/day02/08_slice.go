package main

import "fmt"

func main() {
	fmt.Println("=== 切片（动态数组）===\n")

	// 创建切片的3种方式
	fmt.Println("--- 创建切片 ---")

	// 方式1：从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // 左闭右开[1, 4)，包含索引1,2,3
	fmt.Printf("从数组切片: %v\n", slice1)

	// 方式2：直接声明
	slice2 := []int{10, 20, 30}
	fmt.Printf("直接声明: %v\n", slice2)

	// 方式3：make创建
	slice3 := make([]int, 5)    // 长度5，容量5，初始化为0
	slice4 := make([]int, 3, 5) // 长度3，容量5
	fmt.Printf("make创建(len=5): %v\n", slice3)
	fmt.Printf("make创建(len=3, cap=5): %v\n\n", slice4)

	// 长度和容量
	fmt.Println("--- 长度和容量 ---")
	s := make([]int, 3, 5)
	fmt.Printf("切片: %v\n", s)
	fmt.Printf("长度len: %d, 容量cap: %d\n\n", len(s), cap(s))

	// append：追加元素（重要！）
	fmt.Println("--- append 追加元素 ---")
	nums := []int{1, 2, 3}
	fmt.Printf("原切片: %v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("追加4: %v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	nums = append(nums, 5, 6, 7)
	fmt.Printf("追加5,6,7: %v (len=%d, cap=%d)\n\n", nums, len(nums), cap(nums))

	// 切片扩容机制（当len超过cap时自动扩容）
	fmt.Println("--- 扩容机制 ---")
	s2 := make([]int, 0, 2)
	fmt.Printf("初始: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	s2 = append(s2, 1)
	fmt.Printf("追加1: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	s2 = append(s2, 2)
	fmt.Printf("追加2: %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	s2 = append(s2, 3) // 容量不足，扩容（通常扩容为原来的2倍）
	fmt.Printf("追加3（扩容）: %v (len=%d, cap=%d)\n\n", s2, len(s2), cap(s2))

	// 切片截取
	fmt.Println("--- 切片截取 ---")
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("原切片: %v\n", data)
	fmt.Printf("data[2:5]: %v\n", data[2:5]) // [2 3 4]
	fmt.Printf("data[:5]: %v\n", data[:5])   // [0 1 2 3 4]
	fmt.Printf("data[5:]: %v\n", data[5:])   // [5 6 7 8 9]
	fmt.Printf("data[:]: %v\n\n", data[:])   // 整个切片

	// copy：复制切片
	fmt.Println("--- copy 复制切片 ---")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 100
	fmt.Printf("src: %v\n", src)   // [1 2 3]
	fmt.Printf("dst: %v\n\n", dst) // [100 2 3]

	// 切片是引用类型（重要！）
	fmt.Println("--- 切片是引用类型 ---")
	a := []int{1, 2, 3}
	b := a // 共享底层数组
	b[0] = 100
	fmt.Printf("a: %v\n", a) // [100 2 3]
	fmt.Printf("b: %v\n", b) // [100 2 3]
	fmt.Println("切片是引用类型，修改b会影响a")
}
