func lemonadeChange(bills []int) bool {
    five := 0
    ten := 0

    for i := 0; i < len(bills); i++ {
        if bills[i] == 5 {
            five++
        } else if bills[i] == 10 {
            if five > 0 {
                five--
                ten++
            } else {
                return false
            }
        } else {
            if ten > 0 && five > 0 {
                five--
                ten--
            } else if ten == 0 && five >= 3 {
                five = five - 3
            } else {
                return false
            }
        }
    }

    return true
}