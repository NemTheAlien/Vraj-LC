package main

import "fmt"

func main() {

	/*
		This contains the solution for LeetCode problem 27 - Remove element
	*/

	val1 := removeElement([]int{3, 2, 2, 3}, 3)
	val2 := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)

	fmt.Println(val1, val2)

}

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
