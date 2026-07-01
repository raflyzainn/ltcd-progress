func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0 
    
    for i := 0; i <= len(flowerbed) - 1; i++ {
        if  flowerbed[i] == 0 && (i == 0 || flowerbed[i - 1] == 0) && (i == len(flowerbed) - 1 || flowerbed[i + 1] == 0) {
            flowerbed[i] = 1
            count = count + 1
        } 
    }

    return count >= n
}