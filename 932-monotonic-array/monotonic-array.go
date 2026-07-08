func isMonotonic(nums []int) bool {
    isIncreasing := true 
    isDecreasing := true
    
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] > nums[i + 1] {
            isIncreasing = false 
        } else if nums[i] < nums[i + 1] {
            isDecreasing = false
        }
    }

    if isIncreasing || isDecreasing {
        return true
    }

    return false

}