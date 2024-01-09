func removeElement(nums []int, val int) int {
	counter := 0
	for _, value := range nums {
		if value != val {
			nums[counter] = value
			counter++
		}
	}
	return counter
}
