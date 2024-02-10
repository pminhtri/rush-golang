package solves

func SearchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}

		if i == len(nums)-1 {
			return len(nums)
		}

		if v < target && nums[i+1] > target {

			return i + 1
		}

	}

	return 0
}
