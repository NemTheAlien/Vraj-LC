package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, val := range nums {
		lf := target - val
		if _, ok := seen[lf]; ok {
			return []int{seen[lf], i}
		} else {
			seen[val] = i
		}
	}
	return []int{}
}
