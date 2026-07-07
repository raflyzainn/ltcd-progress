func minCostClimbingStairs(cost []int) int {
    // count := 0
    i := 2; 
    dp := make([]int, len(cost))
    dp[0] = cost[0]  // biaya nyampe tangga 0 = cost[0] sendiri
    dp[1] = cost[1]  // biaya nyampe tangga 1 = cost[1] sendiri (karena boleh start dari sini)

    for i < len(cost) {
        // if cost[i] < cost [i + 1] {
        //     count = count + cost[i]
        //     i = i + 2
        // } else {
        //     count = count + cost[i + 1]
        //     i = i + 1
        // }
        dp[i] = cost[i] + min(dp[i-1], dp[i-2])
        i = i + 1
    }

    return min(dp[len(dp)-1], dp[len(dp)-2])

}