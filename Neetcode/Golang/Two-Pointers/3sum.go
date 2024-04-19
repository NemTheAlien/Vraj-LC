package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for num1Idx := 0; num1Idx < len(nums)-2; num1Idx++ {
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}
		l, r := num1Idx+1, len(nums)-1
		for l < r {
			threeSum := nums[num1Idx] + nums[l] + nums[r]
			if threeSum == 0 {
				res = append(res, []int{nums[num1Idx], nums[l], nums[r]})
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if threeSum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
