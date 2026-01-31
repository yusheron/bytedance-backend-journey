package main

import "fmt"

/*
LeetCode 26 - 删除有序数组中的重复项
难度：简单
标签：数组、双指针

题目描述：
给你一个升序排列的数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
返回删除后数组的新长度。元素的相对顺序应该保持一致。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。
更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

示例：
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4,_,_,_,_,_]
*/

// 双指针法
// 思路：
// - slow指针指向不重复元素的末尾
// - fast指针遍历数组
// - 当nums[fast] != nums[slow]时，将nums[fast]复制到slow+1位置
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// slow指针，指向不重复元素的末尾
	slow := 0

	// fast指针遍历数组
	for fast := 1; fast < len(nums); fast++ {
		// 如果当前元素与slow指向的元素不同
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	// slow+1就是不重复元素的个数
	return slow + 1
}

func main() {
	// 测试用例1
	nums1 := []int{1, 1, 2}
	length1 := removeDuplicates(nums1)
	fmt.Printf("输入: [1,1,2]\n")
	fmt.Printf("输出: %d, nums = %v\n", length1, nums1[:length1])
	// 期望：2, [1,2]

	fmt.Println()

	// 测试用例2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length2 := removeDuplicates(nums2)
	fmt.Printf("输入: [0,0,1,1,1,2,2,3,3,4]\n")
	fmt.Printf("输出: %d, nums = %v\n", length2, nums2[:length2])
	// 期望：5, [0,1,2,3,4]

	fmt.Println()

	// 测试用例3（空数组）
	nums3 := []int{}
	length3 := removeDuplicates(nums3)
	fmt.Printf("输入: []\n")
	fmt.Printf("输出: %d, nums = %v\n", length3, nums3)
	// 期望：0, []

	fmt.Println()

	// 测试用例4（单个元素）
	nums4 := []int{1}
	length4 := removeDuplicates(nums4)
	fmt.Printf("输入: [1]\n")
	fmt.Printf("输出: %d, nums = %v\n", length4, nums4[:length4])
	// 期望：1, [1]
}
