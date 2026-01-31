package main

import "fmt"

func main() {
	fmt.Println("=== for 循环（Go唯一的循环语句）===\n")

	// 形式1：类似C语言的for（初始化; 条件; 后置语句）
	fmt.Println("--- 形式1：标准for循环 ---")
	for i := 1; i <= 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	fmt.Println()

	// 形式2：类似while循环（只有条件）
	fmt.Println("--- 形式2：类似while ---")
	j := 1
	for j <= 5 {
		fmt.Printf("j = %d\n", j)
		j++
	}
	fmt.Println()

	// 形式3：无限循环
	fmt.Println("--- 形式3：无限循环（带break）---")
	k := 1
	for {
		if k > 5 {
			break // 跳出循环
		}
		fmt.Printf("k = %d\n", k)
		k++
	}
	fmt.Println()

	// break和continue
	fmt.Println("--- break 和 continue ---")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue // 跳过本次循环
		}
		if i == 8 {
			break // 跳出整个循环
		}
		fmt.Printf("i = %d\n", i)
	}
	fmt.Println()

	// 嵌套循环
	fmt.Println("--- 嵌套循环：打印乘法表 ---")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()

	// range遍历（后面学数组、切片、map时重点使用）
	fmt.Println("--- range 遍历数组 ---")
	nums := [5]int{10, 20, 30, 40, 50}
	for index, value := range nums {
		fmt.Printf("索引%d: 值%d\n", index, value)
	}
	fmt.Println()

	// 只要索引
	for index := range nums {
		fmt.Printf("索引%d\n", index)
	}
	fmt.Println()

	// 只要值（用_忽略索引）
	for _, value := range nums {
		fmt.Printf("值%d\n", value)
	}
}
