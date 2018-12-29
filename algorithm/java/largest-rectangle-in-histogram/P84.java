func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    maxCount := 0
    for start, _ := range heights {
        cur := heights[start]
        count := 1
        for p := start-1; p >= 0; p-- {
            if heights[p] >= cur {
                count++
            } else {
                break
            }
        }
        for p := start+1; p < len(heights); p++ {
            if heights[p] >= cur {
                count++
            } else {
                break
            }
        }
        if count * cur > maxCount {
            maxCount = count * cur
        }
    }
    return maxCount
}
