func findDisappearedNumbers(nums []int) []int {
    m := make(map[int]int)
    var result []int

    for i := 0; i < len(nums); i++ {
        m[nums[i]]++
    }

    for i := 1; i <= len(nums); i++ {
        if m[i] == 0 {
            result = append(result, i)
        }
    }

    return result
}