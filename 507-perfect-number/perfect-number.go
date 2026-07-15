func checkPerfectNumber(num int) bool {
    jumlah := 0 

    for i := 1; i <= num/2; i++ {  // gak perlu sampe num, divisor max itu num/2
        if num % i == 0 {
            jumlah = jumlah + i
        }
        if jumlah > num {  // udah kelewat, gak mungkin balik pas lagi
            return false
        }
    }

    return jumlah == num
}