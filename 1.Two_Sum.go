package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Print(result)
}

func twoSum(nums []int, target int) []int {
	for j := 0; j < len(nums); j++ {
		for i := j + 1; i < len(nums); i++ {
			if nums[j]+nums[i] == target {
				return []int{j, i}
			}
		}
	}
	return nil
}
