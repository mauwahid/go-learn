package code_pattern

import "sort"

func findSumOfThree(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			triple := nums[i] + nums[left] + nums[right]

			if triple == target {
				return true
			} else if triple < target {
				left++
			} else {
				right--
			}

		}
	}

	return false
}

func findSumOfThree2(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))

	low, high, triple := 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {

		low = i + 1
		high = len(nums) - 1

		for low < high {
			triple = nums[i] + nums[low] + nums[high]

			if triple == target {
				return true
			} else if triple < target {
				low += 1
			} else {
				high -= 1
			}
		}
	}
	return false
}
