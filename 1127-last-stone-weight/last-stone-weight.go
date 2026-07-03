func deleteAt(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}

func lastStoneWeight(stones []int) int {
    if len(stones) == 1 {
        return stones[0]
    }
    
    for len(stones) > 1 {
        largest := 0
        second_largest := -1

        for i := 0; i < len(stones); i++ {
            if stones[i] > stones[largest] {
                second_largest = largest
                largest = i 
            } else if i != largest && (second_largest == -1 || (stones[i] > stones[second_largest] && stones[i] <= stones[largest])) {
                second_largest = i
            }
        }

        if stones[largest] != stones[second_largest] {
            batuBaru := stones[largest] - stones[second_largest]

            if largest > second_largest {
                stones = deleteAt(stones, largest)
                stones = deleteAt(stones, second_largest)
            } else {
                stones = deleteAt(stones, second_largest)
                stones = deleteAt(stones, largest)
            }

            stones = append(stones, batuBaru)
        } else if stones[largest] == stones[second_largest] {
            if largest > second_largest {
                stones = deleteAt(stones, largest)
                stones = deleteAt(stones, second_largest)
            } else {
                stones = deleteAt(stones, second_largest)
                stones = deleteAt(stones, largest)
            }
        }
    }

    if len(stones) == 0 {
        return 0
    }

    for i := 0; i < len(stones); i++ {
        return stones[i]
    }

    return stones[0]
}