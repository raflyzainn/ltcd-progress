func distributeCandies(candyType []int) int {
    m := make(map[int]int)
    n := len(candyType)
    count := 0

    for i := 0; i < n; i++ {
        m[candyType[i]]++
    }

    for key, value := range m {
        fmt.Println(key, value)      
        count++
    }

    fmt.Println("count", count)
    fmt.Println("n", n)

    if count == n/2 {
        return count 
    } else if count > n/2 {
        return n / 2
    }

    return count
}