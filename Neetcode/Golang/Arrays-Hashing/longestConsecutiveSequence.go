package main

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, val := range nums {
		set[val] = true
	}
	res := 0
	for _, val := range nums {
		if set[val-1] {
			continue
		}
		sequence := 1
		temp := val + 1
		for set[temp] {
			sequence++
			temp++
		}
		res = max(res, sequence)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
