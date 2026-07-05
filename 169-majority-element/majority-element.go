func majorityElement(nums []int) int {
    m := make(map[int]int)
    majority := len(nums) / 2

    for i := 0; i < len(nums); i++ {
        m[nums[i]]++

        for key, value := range m {
            if value > majority {
                return key 
            }
        }
    }

    

    return 0
}