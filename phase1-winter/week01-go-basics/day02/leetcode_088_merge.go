package main

import "fmt"

/*
LeetCode 88 - 合并两个有序数组
难度：简单
标签：数组、双指针

题目描述：
给你两个按非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，
分别表示 nums1 和 nums2 中的元素数目。
请你合并 nums2 到 nums1 中，使合并后的数组同样按非递减顺序排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，
后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
*/

// 解法1：双指针（从后往前）- 推荐
// 思路：从后往前填充，避免覆盖nums1的元素
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	// p1指向nums1的有效元素末尾
	p1 := m - 1
	// p2指向nums2的末尾
	p2 := n - 1
	// p指向nums1的末尾（填充位置）
	p := m + n - 1

	// 从后往前比较并填充
	for p2 >= 0 {
		// 如果p1已经用完，直接复制nums2剩余元素
		if p1 < 0 {
			nums1[p] = nums2[p2]
			p2--
		} else if nums1[p1] > nums2[p2] {
			// nums1[p1]更大，放到p位置
			nums1[p] = nums1[p1]
			p1--
		} else {
			// nums2[p2]更大或相等，放到p位置
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
}

// 解法2：暴力法（先合并后排序）- 不推荐
// 时间复杂度：O((m+n)log(m+n))
// 空间复杂度：O(log(m+n))
func mergeBruteForce(nums1 []int, m int, nums2 []int, n int) {
	// 将nums2复制到nums1的后面
	copy(nums1[m:], nums2)

	// 使用冒泡排序（为了练习，实际可用sort.Ints）
	for i := 0; i < m+n-1; i++ {
		for j := 0; j < m+n-1-i; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
}

func main() {
	// 测试用例1
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	fmt.Printf("合并前: nums1=%v, nums2=%v\n", nums1, nums2)
	merge(nums1, m, nums2, n)
	fmt.Printf("合并后: %v\n", nums1)
	// 期望输出：[1, 2, 2, 3, 5, 6]

	fmt.Println()

	// 测试用例2
	nums3 := []int{1}
	nums4 := []int{}
	m2, n2 := 1, 0
	fmt.Printf("合并前: nums1=%v, nums2=%v\n", nums3, nums4)
	merge(nums3, m2, nums4, n2)
	fmt.Printf("合并后: %v\n", nums3)
	// 期望输出：[1]

	fmt.Println()

	// 测试用例3
	nums5 := []int{0}
	nums6 := []int{1}
	m3, n3 := 0, 1
	fmt.Printf("合并前: nums1=%v, nums2=%v\n", nums5, nums6)
	merge(nums5, m3, nums6, n3)
	fmt.Printf("合并后: %v\n", nums5)
	// 期望输出：[1]
}
