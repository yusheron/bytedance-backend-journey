package main

func removeDuplicates(nums []int) int {
	nums2 := make([]int, 0)
	x := 1
	y := 0
	nums2[0] = nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums2[x] = nums[i]
			x++
		} else {
			y++
		}
		for y > 0 {
			nums2 = append(nums2, 0)
			y--
		}
	}
	return x
}
