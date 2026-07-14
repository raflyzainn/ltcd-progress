func sortPeople(names []string, heights []int) []string {
    n := len(heights)
    
    // bikin index awal [0, 1, 2, ..., n-1]
    index := make([]int, n)
    for i := 0; i < n; i++ {
        index[i] = i
    }

    // sorting index berdasarkan heights, descending
    for i := 0; i < n; i++ {
        for j := 0; j < n-1; j++ {
            if heights[index[j]] < heights[index[j+1]] {
                index[j], index[j+1] = index[j+1], index[j]
            }
        }
    }

    // ambil nama sesuai urutan index yang udah disortir
    result := make([]string, n)
    for k := 0; k < n; k++ {
        result[k] = names[index[k]]
    }

    return result
}