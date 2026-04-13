func longestCommonPrefix(strs []string) string {
    for i := 0; i < len(strs[0]); i++ {
        char := strs[0][i]

        for _, v := range strs {
            if i >= len(v) || v[i] != char {
                charv := strs[0][:i]
                return charv
            } 
        }
    }

    return strs[0]
} 