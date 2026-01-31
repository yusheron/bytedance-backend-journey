package main

import "fmt"

func main() {
	fmt.Println("=== if 条件语句 ===\n")

	// 基本if
	score := 85
	if score >= 60 {
		fmt.Printf("分数%d: 及格\n\n", score)
	}

	// if-else
	age := 17
	if age >= 18 {
		fmt.Println("你已成年")
	} else {
		fmt.Println("你未成年")
	}
	fmt.Println()

	// if-else if-else
	grade := 92
	if grade >= 90 {
		fmt.Printf("成绩%d: 优秀\n", grade)
	} else if grade >= 80 {
		fmt.Printf("成绩%d: 良好\n", grade)
	} else if grade >= 60 {
		fmt.Printf("成绩%d: 及格\n", grade)
	} else {
		fmt.Printf("成绩%d: 不及格\n", grade)
	}
	fmt.Println()

	// if的简短语句（重要特性！）
	// 在条件判断前执行一个简单语句，变量作用域仅在if-else块内
	if num := 10; num%2 == 0 {
		fmt.Printf("%d 是偶数\n", num)
	} else {
		fmt.Printf("%d 是奇数\n", num)
	}
	// fmt.Println(num) // 这里访问不到num，会报错
	fmt.Println()

	// 实际应用场景：错误处理
	if value := getValue(); value > 0 {
		fmt.Printf("获取到的值: %d (正数)\n", value)
	} else if value == 0 {
		fmt.Println("获取到的值: 0")
	} else {
		fmt.Printf("获取到的值: %d (负数)\n", value)
	}
}

func getValue() int {
	return 42
}
