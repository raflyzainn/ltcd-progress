func thirdMax(nums []int) int {
    sort.Ints(nums)
	
	distinctCount := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 || nums[i] != nums[i+1] {
			distinctCount++
			if distinctCount == 3 {
				return nums[i]
			}
		}
	}

	return nums[len(nums)-1]
}