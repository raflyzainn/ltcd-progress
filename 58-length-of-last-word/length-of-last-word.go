func lengthOfLastWord(s string) int {
    count := 0

    i := len(s) - 1

    for i >= 0 && s[i] == ' ' {
        i--
    }

    for i >= 0 && s[i] != ' ' {
        count++
        i--
    }
    
    // for i := len(s) - 1; i >= 0; i-- {
    //     if s[i] == ' ' {
    //         continue
    //     } else if s[i] != ' '{
    //         break
    //     } else {
    //         count += 1
    //     }
    // }

    return count
}