func heightChecker(heights []int) int {
    expected := make([]int, len(heights))
    copy(expected, heights)
    
    for i := 0; i < len(expected); i++ {
        for j := i + 1; j < len(expected); j++ {
            if expected[i] > expected[j] {
                temp := expected[i];
                expected[i] = expected[j];
                expected[j] = temp;
            }
        }
    }

    count := 0

    for i := 0; i < len(heights); i++ {
        if expected[i] != heights[i] {
            count = count + 1
        }
    }

    return count
}