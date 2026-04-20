func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := make([]int, 26)

    for i := 0; i < len(s); i++ {
        index := s[i] - 'a'
        count[index]++
    }

    for i := 0; i < len(s); i++ {
        index := t[i] - 'a'
        count[index]--
    }

    for i := 0; i < 26; i++ {
        if count[i] != 0 {
            return false
        } 
    }

    return true
}