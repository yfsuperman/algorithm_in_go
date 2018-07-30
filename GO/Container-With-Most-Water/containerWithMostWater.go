package containerWithMostWater

// Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

// Example:

// Input: [1,8,6,2,5,4,8,3,7]
// Output: 49

func maxArea(height []int) int {
    i, j, area := 0, len(height) - 1, 0
    for i < j {
        area = getMax(area, getMin(height[i], height[j]) * (j - i))
        if height[i] > height[j] {
            j --
        } else {
            i ++
        }
    }
    
    return area
}

func getMin(a int, b int) int {
    if a > b {
        return b
    }
    return a
}

func getMax(a int, b int) int {
    if a < b {
        return b
    }
    return a
}
