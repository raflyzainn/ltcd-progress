func reverse(x int) int {
    result := 0
    for x != 0 {
        digit := x % 10
        x /= 10
        result = result*10 + digit
        if result > 2147483647 {
            return 0 
        } else if result < -2147483648 {
            return 0
        }
    }
    return result
}