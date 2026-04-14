func isValid(s string) bool {
    stack := []rune{}

    for i := 0; i < len(s); i++ {
        char := rune(s[i])
        if char == '(' || char == '{' || char == '['{
            stack = append(stack, rune(s[i]))
        } else {
            if len(stack) == 0 {
                return false
            }
            n := len(stack) - 1
            top := stack[n]
            if char == ')' && top != '(' {
                return false
            } else if char == '}' && top != '{' {
                return false
            } else if char == ']' && top != '[' {
                return false
            }
            stack = stack[:n]
        }
    }
    return len(stack) == 0
}