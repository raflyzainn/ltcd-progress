func arrangeCoins(n int) int {
    if n == 1 {
        return 1
    }
    
    count := 1
    
    for n > 0 {
        n = n - count
        
        if n < 0 {
            return count - 1
        } else if n == 0 {
            return count
        } else {
            count = count + 1
        }
    }

    return 0
}