func moveZeroes(nums []int)  {
    if (len(nums) == 0) {
        fmt.Println(nums)
    }

    insertPos := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[insertPos] = nums[i]
            insertPos++
        }
    }

    for j := insertPos; j < len(nums); j++ {
        nums[j] = 0
    }

    fmt.Println(nums)
}