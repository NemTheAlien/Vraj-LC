func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[counter] != nums[i] {
			counter++
			nums[counter] = nums[i]
		}
	}
	return counter + 1
}