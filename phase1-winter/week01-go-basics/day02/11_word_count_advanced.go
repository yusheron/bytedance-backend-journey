package main

import (
	"fmt"
	"sort"
	"strings"
)

// WordFreq 词频结构体
type WordFreq struct {
	word  string
	count int
}

func main() {
	fmt.Println("=== 词频统计程序（增强版）===\n")

	// 示例文本
	text := `
	ByteDance is a great company.
	ByteDance develops TikTok and Douyin.
	I want to join ByteDance as a backend developer.
	Go is the main language at ByteDance.
	ByteDance ByteDance ByteDance!
	`

	// 步骤1：转小写、分割单词、过滤标点
	text = strings.ToLower(text) // 转小写
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "!", "")

	words := strings.Fields(text) // Fields按空白字符分割（比Split更智能）
	fmt.Printf("单词总数: %d\n", len(words))

	// 步骤2：统计词频
	wordCount := make(map[string]int)
	for _, word := range words {
		if word != "" { // 过滤空字符串
			wordCount[word]++
		}
	}

	// 步骤3：转换为切片以便排序
	var wordFreqs []WordFreq
	for word, count := range wordCount {
		wordFreqs = append(wordFreqs, WordFreq{word, count})
	}

	// 步骤4：按出现次数排序（从高到低）
	sort.Slice(wordFreqs, func(i, j int) bool {
		return wordFreqs[i].count > wordFreqs[j].count
	})

	// 步骤5：打印结果（前10个高频词）
	fmt.Println("\n=== 词频统计结果（按频率降序）===")
	limit := 10
	if len(wordFreqs) < limit {
		limit = len(wordFreqs)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%2d. %-15s: %d次\n", i+1, wordFreqs[i].word, wordFreqs[i].count)
	}

	// 额外功能：统计信息
	fmt.Println("\n=== 统计信息 ===")
	fmt.Printf("不同单词数: %d\n", len(wordCount))
	fmt.Printf("总单词数: %d\n", len(words))
}
