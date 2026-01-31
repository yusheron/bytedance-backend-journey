package main

import "fmt"

func main() {
	fmt.Println("=== Go 常量 ===\n")

	// 普通常量
	const PI = 3.14159
	const CompanyName = "字节跳动"
	fmt.Printf("PI = %f\n", PI)
	fmt.Printf("公司名称 = %s\n\n", CompanyName)

	// const不能用 := 声明
	// const name := "error" // 这会报错

	// iota：常量计数器
	fmt.Println("=== iota 常量计数器 ===")

	// 示例1：枚举星期
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Printf("星期日=%d, 星期一=%d, 星期六=%d\n\n", Sunday, Monday, Saturday)

	// 示例2：文件权限（使用位移）
	const (
		ReadPermission    = 1 << iota // 1 << 0 = 1 (二进制：001)
		WritePermission               // 1 << 1 = 2 (二进制：010)
		ExecutePermission             // 1 << 2 = 4 (二进制：100)
	)
	fmt.Printf("读权限=%d, 写权限=%d, 执行权限=%d\n", ReadPermission, WritePermission, ExecutePermission)

	// 组合权限（位运算）
	fullPermission := ReadPermission | WritePermission | ExecutePermission
	fmt.Printf("完整权限=%d (二进制：111)\n\n", fullPermission)

	// 示例3：跳过某些值
	const (
		a = iota // 0
		b        // 1
		_        // 2（跳过）
		c        // 3
	)
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
}
