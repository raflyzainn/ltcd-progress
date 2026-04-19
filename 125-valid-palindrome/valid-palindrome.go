func isPalindrome(s string) bool {
    if s == "" {
        return true
    }

    b := []byte(s)
    
    for i := 0; i < len(b); i++ {
        // Check if character is within the uppercase ASCII range (A-Z)
        if b[i] >= 'A' && b[i] <= 'Z' {
            // Add 32 to shift it to the lowercase range
            b[i] += 32
        }
    }

    lowerCaseString := string(b)

    left := 0
    right := len(lowerCaseString) - 1

    for left < right {
        bukanHurufLeft := !(lowerCaseString[left] >= 'a' && lowerCaseString[left] <= 'z') 
        bukanAngkaLeft := !(lowerCaseString[left] >= '0' && lowerCaseString[left] <= '9') 
        bukanHurufRight := !(lowerCaseString[right] >= 'a' && lowerCaseString[right] <= 'z') 
        bukanAngkaRight := !(lowerCaseString[right] >= '0' && lowerCaseString[right] <= '9') 
        if bukanHurufLeft && bukanAngkaLeft {
            left++
        } else if bukanHurufRight && bukanAngkaRight{
            right--
        } else {
            if lowerCaseString[left] != lowerCaseString[right] {
                return false
            } else {
                left++
                right--
            }
        }
    }

    return true
}