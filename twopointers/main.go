package main

import (
	"fmt"
	"slices"
)

func twoPointers(nums []int, target int) *int {
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return &sum
		} else if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return nil
}

func main() {
	fmt.Println(*twoPointers([]int{1, 2, 3}, 3))
	fmt.Println(*twoPointers([]int{1, 2, 3}, 2))
	fmt.Println(*twoPointers([]int{1, 2, 3}, 1))
	fmt.Println(*twoPointers([]int{6, 2, 3, 9, 10, 4, 4}, 12))
}
