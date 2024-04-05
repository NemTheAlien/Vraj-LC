package main

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	countElem := make([][]int, len(nums)+1)
	var res []int

	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = 1
		} else {
			seen[num] += 1
		}
	}

	for key, val := range seen {
		countElem[val] = append(countElem[val], key)
	}

	for i := len(countElem) - 1; i > 0; i-- {
		res = append(res, countElem[i]...)
		if len(res) == k {
			return res
		}
	}

	return res
}
