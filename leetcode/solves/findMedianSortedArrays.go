package solves

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))

	i, j, k := 0, 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) {
		nums[k] = nums1[i]
		i++
		k++
	}

	for j < len(nums2) {
		nums[k] = nums2[j]
		j++
		k++
	}

	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
	} else {
		return float64(nums[len(nums)/2])
	}
}
