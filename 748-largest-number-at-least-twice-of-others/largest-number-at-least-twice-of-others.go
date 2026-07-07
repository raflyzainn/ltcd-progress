func dominantIndex(nums []int) int {
    maxIndex := 0
    secondLargestNumber := 0
    maxNumber := 0

    for i := 0; i < len(nums); i++ {
        if maxNumber < nums[i]  {
            secondLargestNumber = maxNumber
            maxNumber = nums[i]
            maxIndex = i
            
        }  else if nums[i] > secondLargestNumber && secondLargestNumber != maxNumber  {
            secondLargestNumber = nums[i]
        }
    }

    fmt.Println(maxNumber, maxIndex, secondLargestNumber)

    if secondLargestNumber * 2 > maxNumber {
        return -1
    } else {
        return maxIndex
    }

    return 0
}