package main

import (
	"fmt"
	"slices"
)

func twoPointers(nums []int, target int) []int {
	slices.Sort(nums)
	n := len(nums)
	r := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				return []int{left, right}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return r
}

func main() {
	fmt.Println(twoPointers([]int{1, 2, 3}, 3))
	fmt.Println(twoPointers([]int{1, 2, 3}, 2))
	fmt.Println(twoPointers([]int{1, 2, 3}, 1))
	fmt.Println(twoPointers([]int{6, 2, 3, 9, 10, 4, 4}, 12))
}
