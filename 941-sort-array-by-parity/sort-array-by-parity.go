func sortArrayByParity(nums []int) []int {
    // flag := 0; 
    simpen := 0
    
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 == 0 {
            // nums[flag] = nums[i]
            for j := 0; j < len(nums); j++ {
                if nums[j] % 2 != 0 {
                    simpen = nums[j]
                    nums[j] = nums[i]
                    nums[i] = simpen
                }
            }

        }
    }

    fmt.Println(nums)

    return nums
}