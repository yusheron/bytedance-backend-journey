package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== switch 语句 ===\n")

	// 基本switch
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6, 7: // 多个条件
		fmt.Println("周末")
	default:
		fmt.Println("无效的天数")
	}
	fmt.Println()

	// switch的简短语句
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Printf("现在是上午 %d 点\n", hour)
	case hour < 18:
		fmt.Printf("现在是下午 %d 点\n", hour)
	default:
		fmt.Printf("现在是晚上 %d 点\n", hour)
	}
	fmt.Println()

	// 无条件的switch（相当于if-else if-else）
	score := 85
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
	fmt.Println()

	// switch类型判断（后续学接口时会用到）
	var x interface{} = 100
	switch x.(type) {
	case int:
		fmt.Println("x 是 int 类型")
	case string:
		fmt.Println("x 是 string 类型")
	default:
		fmt.Println("x 是其他类型")
	}
	fmt.Println()

	// fallthrough：强制执行下一个case（少用）
	num := 2
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough // 会继续执行下一个case
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}
	// 输出：two 和 three
}
