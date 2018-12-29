func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    leftLen := make([]int, len(heights))
    rightLen := make([]int, len(heights))
    leftLen[0] = -1
    rightLen[len(heights)-1] = len(heights)

    for i := 1; i < len(heights); i++ {
        p := i-1  
        for  p >= 0 &&  heights[p] >= heights[i] {
            p = leftLen[p]
        }
        leftLen[i] =  p
    }

    for i := len(heights)-2; i >= 0; i-- {
        p := i+1
        for  p < len(heights) && heights[p] >= heights[i]  {
            p = rightLen[p]
        }
        rightLen[i] = p
    }

    fmt.Println(leftLen)
    
    fmt.Println(rightLen)
    maxCount := 0
    for i := 0; i < len(heights); i++ {
        maxCount = max(maxCount,  heights[i] * (rightLen[i] - leftLen[i] - 1))
    }
    return maxCount
}

func max( i ,  j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
