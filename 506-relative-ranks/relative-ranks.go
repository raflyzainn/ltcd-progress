func findRelativeRanks(score []int) []string {
    n := len(score)
    pairs := make([][2]int, n)
    
    for i := 0; i < n; i++ {
        pairs[i] = [2]int{score[i], i}
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if pairs[j][0] < pairs[j+1][0] { 
                pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
            }
        }
    }

    result := make([]string, n)

    for k := 0; k < n; k++ {
        originalIndex := pairs[k][1]
        
        if k == 0 {
            result[originalIndex] = "Gold Medal"
        } else if k == 1 {
            result[originalIndex] = "Silver Medal"
        } else if k == 2 {
            result[originalIndex] = "Bronze Medal"
        } else {
            result[originalIndex] = strconv.FormatInt(int64(k+1), 10)
        }
    }

    fmt.Println(result)

    return result
}