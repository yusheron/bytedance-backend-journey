package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== 词频统计程序（简单版）===\n")

	// 示例文本
	text := "Go is great Go is fast Go is fun Go is simple"

	// 步骤1：分割字符串为单词
	words := strings.Split(text, " ")
	fmt.Printf("分割后的单词: %v\n\n", words)

	// 步骤2：创建map存储词频
	wordCount := make(map[string]int)

	// 步骤3：统计每个单词出现次数
	for _, word := range words {
		wordCount[word]++
	}

	// 步骤4：打印结果
	fmt.Println("词频统计结果:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d次\n", word, count)
	}
}
