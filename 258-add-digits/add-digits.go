func addDigits(num int) int {
    firstDigit := num / 10 
    lastDigit := num % 10 

    total := firstDigit + lastDigit 

    for total >= 10 {
        digit := total / 10 
        num := total % 10 

        total =  digit + num 
    }

    return total 
}