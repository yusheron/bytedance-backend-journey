package main

import "fmt"

// isValid 有效的括号
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func isValid(s string) bool {
    // 奇数长度一定无效
    if len(s)%2 == 1 {
        return false
    }

    // 括号匹配映射：右括号 -> 左括号
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    // 用 slice 模拟栈
    stack := []rune{}

    for _, char := range s {
        // 如果是右括号
        if pair, ok := pairs[char]; ok {
            // 栈为空或栈顶不匹配
            if len(stack) == 0 || stack[len(stack)-1] != pair {
                return false
            }
            // 匹配成功，弹出栈顶
            stack = stack[:len(stack)-1]
        } else {
            // 左括号，入栈
            stack = append(stack, char)
        }
    }

    // 栈为空则所有括号都匹配
    return len(stack) == 0
}

func main() {
    // 测试用例
    testCases := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
    for _, s := range testCases {
        fmt.Printf("s = \"%s\", isValid = %v\n", s, isValid(s))
    }
}