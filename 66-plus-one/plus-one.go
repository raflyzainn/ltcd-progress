func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    // All digits were 9 (e.g. [9,9,9] → [1,0,0,0])
    return append([]int{1}, digits...)
}