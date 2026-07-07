func maximumProduct(nums []int) int {
    n := len(nums) - 1
    sort.Ints(nums)

    fmt.Println(nums)

    option1 := nums[n] * nums[n-1] * nums[n-2]
    option2 := nums[0] * nums[1] * nums[n]

    if option1 > option2 {
        return option1
    }
    return option2
}