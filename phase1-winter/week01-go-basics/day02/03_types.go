package main

import "fmt"

func main() {
	fmt.Println("=== Go 基本数据类型 ===\n")

	// 1. 整型
	fmt.Println("--- 整型 ---")
	var i8 int8 = 127   // -128 到 127
	var ui8 uint8 = 255 // 0 到 255
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	fmt.Printf("int8: %d, uint8: %d\n", i8, ui8)
	fmt.Printf("int32: %d, int64: %d\n\n", i32, i64)

	// int类型（根据平台自动选择32位或64位）
	var num int = 100
	fmt.Printf("int: %d, 类型: %T\n\n", num, num)

	// 2. 浮点型
	fmt.Println("--- 浮点型 ---")
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	fmt.Printf("float32: %f\n", f32)
	fmt.Printf("float64: %.15f\n\n", f64) // 保留15位小数

	// 3. 布尔型
	fmt.Println("--- 布尔型 ---")
	var isStudent bool = true
	var hasJob bool = false
	fmt.Printf("isStudent: %v, hasJob: %v\n\n", isStudent, hasJob)

	// 4. 字符串
	fmt.Println("--- 字符串 ---")
	var str1 string = "Hello, 字节跳动!"
	str2 := "Go语言学习"
	fmt.Printf("str1: %s\n", str1)
	fmt.Printf("str2: %s\n", str2)
	fmt.Printf("字符串长度: %d\n", len(str2)) // len()返回字节数，不是字符数
	fmt.Printf("字符串拼接: %s\n\n", str1+" "+str2)

	// 5. 字符（rune和byte）
	fmt.Println("--- 字符 ---")
	var ch1 rune = '中' // rune是int32的别名，用于表示Unicode字符
	var ch2 byte = 'A' // byte是uint8的别名，用于表示ASCII字符
	fmt.Printf("rune字符: %c (Unicode: %d)\n", ch1, ch1)
	fmt.Printf("byte字符: %c (ASCII: %d)\n\n", ch2, ch2)

	// 6. 类型转换（Go不支持隐式转换，必须显式转换）
	fmt.Println("--- 类型转换 ---")
	var intNum int = 100
	var floatNum float64 = float64(intNum) // int转float64
	fmt.Printf("int转float64: %d -> %f\n", intNum, floatNum)

	var pi float64 = 3.14
	var intPi int = int(pi) // float64转int（会截断小数部分）
	fmt.Printf("float64转int: %f -> %d\n\n", pi, intPi)

	// 字符串和数字转换（需要用strconv包，这里先了解概念）
	fmt.Println("--- 注意 ---")
	fmt.Println("字符串和数字的转换需要用 strconv 包，稍后会学习")
}
