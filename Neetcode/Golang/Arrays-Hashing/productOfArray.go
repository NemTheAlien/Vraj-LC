package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i, val := range nums {
		res[i] = prefix
		prefix *= val
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
