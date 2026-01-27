package main

import "fmt"

// twoSum 两数之和
// 题目：给定一个整数数组 nums 和一个目标值 target，找出数组中和为目标值的两个数的下标
// 思路：用哈希表存储已遍历的数，查找 target - num 是否存在
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum(nums []int, target int) []int {
    // 哈希表：存储 值 -> 下标
    hashMap := make(map[int]int)

    for i, num := range nums {
        // 计算需要的另一个数
        complement := target - num

        // 如果哈希表中存在，直接返回
        if j, ok := hashMap[complement]; ok {
            return []int{j, i}
        }

        // 否则将当前数存入哈希表
        hashMap[num] = i
    }

    return nil
}

func main() {
    // 测试用例 1
    nums1 := []int{2, 7, 11, 15}
    target1 := 9
    result1 := twoSum(nums1, target1)
    fmt.Printf("输入: nums = %v, target = %d\n", nums1, target1)
    fmt.Printf("输出: %v\n\n", result1) // [0, 1]

    // 测试用例 2
    nums2 := []int{3, 2, 4}
    target2 := 6
    result2 := twoSum(nums2, target2)
    fmt.Printf("输入: nums = %v, target = %d\n", nums2, target2)
    fmt.Printf("输出: %v\n\n", result2) // [1, 2]

    // 测试用例 3
    nums3 := []int{3, 3}
    target3 := 6
    result3 := twoSum(nums3, target3)
    fmt.Printf("输入: nums = %v, target = %d\n", nums3, target3)
    fmt.Printf("输出: %v\n", result3) // [0, 1]
}