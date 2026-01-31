package main

import "fmt"

func main() {
	fmt.Println("=== map（字典/映射）===\n")

	// 创建map的3种方式
	fmt.Println("--- 创建map ---")

	// 方式1：make创建
	map1 := make(map[string]int)
	fmt.Printf("make创建: %v\n", map1)

	// 方式2：字面量初始化
	map2 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 7,
	}
	fmt.Printf("字面量初始化: %v\n\n", map2)

	// 方式3：声明（但未初始化，值为nil，不能直接使用）
	var map3 map[string]int
	fmt.Printf("未初始化: %v\n", map3)
	// map3["key"] = 1 // panic: assignment to entry in nil map
	fmt.Println()

	// 增、改
	fmt.Println("--- 增加和修改元素 ---")
	scores := make(map[string]int)
	scores["Alice"] = 95 // 增加
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("添加后: %v\n", scores)

	scores["Alice"] = 98 // 修改
	fmt.Printf("修改后: %v\n\n", scores)

	// 查
	fmt.Println("--- 查询元素 ---")
	score, exists := scores["Alice"]
	if exists {
		fmt.Printf("Alice的分数: %d\n", score)
	}

	score2, exists2 := scores["David"]
	if !exists2 {
		fmt.Printf("David不存在，默认值: %d\n\n", score2) // 0
	}

	// 删
	fmt.Println("--- 删除元素 ---")
	fmt.Printf("删除前: %v\n", scores)
	delete(scores, "Bob")
	fmt.Printf("删除Bob后: %v\n\n", scores)

	// 遍历map
	fmt.Println("--- 遍历map ---")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}
	fmt.Println()

	// 只遍历key
	for name := range scores {
		fmt.Printf("学生: %s\n", name)
	}
	fmt.Println()

	// map的长度
	fmt.Printf("map长度: %d\n\n", len(scores))

	// map是引用类型
	fmt.Println("--- map是引用类型 ---")
	m1 := map[string]int{"a": 1}
	m2 := m1
	m2["a"] = 100
	fmt.Printf("m1: %v\n", m1) // map[a:100]
	fmt.Printf("m2: %v\n", m2) // map[a:100]

	// 注意：map是无序的
	fmt.Println("\n注意：map遍历是无序的，每次运行顺序可能不同")
}
