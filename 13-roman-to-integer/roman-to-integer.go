func romanToInt(s string) int {
    romanValues := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    result := 0
    n := len(s)
    
    for i := 0; i < n; i++ {
        curr := romanValues[s[i]]
        next := 0
        if i+1 < n {
            next = romanValues[s[i+1]]
        }
        
        if curr < next {
            result -= curr
        } else {
            result += curr
        }
    }
    
    return result
}