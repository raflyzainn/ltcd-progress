func missingNumber(nums []int) int {
    // if len(nums) == 1 {
    //     return 0
    // }
    
    n := len(nums) 

    sort.Ints(nums)

    for i := 0; i < n; i++ {
        if i != nums[i] {
            return i
        }
    }

    return len(nums)
}   