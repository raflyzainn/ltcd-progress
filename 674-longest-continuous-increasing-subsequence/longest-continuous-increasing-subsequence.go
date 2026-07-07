func findLengthOfLCIS(nums []int) int {
    count := 1
    maxCount := 1

    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] < nums[i + 1] {
            count++
        } else {
            if count > maxCount {
                maxCount = count
            }
            count = 1
        }
    }

    if count > maxCount {
        return count 
    } else {
        return maxCount
    }
        
    return 0
}