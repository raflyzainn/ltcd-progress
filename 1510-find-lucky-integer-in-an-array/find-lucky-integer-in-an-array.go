func findLucky(arr []int) int {
    count := make(map[int]int)
    result := -1
    
    for i := 0; i < len(arr); i++ {
        count[arr[i]]++
    }

    for key, value := range count {
        if key == value {
            if key > result {
                result = key
            }
        } 
    }

    return result
}