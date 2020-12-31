package leetcode

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		temp := target - nums[i]

		if _, ok := maps[temp]; ok {
			return []int{maps[temp], i}
		}

		maps[nums[i]] = i

	}

	return nil
}
