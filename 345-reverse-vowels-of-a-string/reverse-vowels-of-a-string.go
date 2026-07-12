func reverseVowels(s string) string {
    b := []byte(s)
    left, right := 0, len(b)-1

    isVowel := func(c byte) bool {
        return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
            c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
    }

    for left < right {
        if !isVowel(b[left]) {
            left++
            continue
        }
        if !isVowel(b[right]) {
            right--
            continue
        }
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }

    return string(b)
}