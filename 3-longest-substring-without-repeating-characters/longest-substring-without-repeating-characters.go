func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLen := 0
    left := 0

    for right := 0; right < len(s); right++ {
        ch := s[right]
        
        // If char was seen and is within our current window, shrink from left
        if idx, found := charIndex[ch]; found && idx >= left {
            left = idx + 1
        }
        
        charIndex[ch] = right
        
        if curr := right - left + 1; curr > maxLen {
            maxLen = curr
        }
    }
    
    return maxLen
}