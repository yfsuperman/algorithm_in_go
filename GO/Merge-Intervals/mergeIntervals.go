package mergeIntervals

// Given a collection of intervals, merge all overlapping intervals.

// Example 1:

// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

import (
	"sort"
)

// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <=1 {
        return intervals
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    
    res := make([]Interval, 0)
    prev := intervals[0]
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start > prev.End {
            res = append(res, prev)
            prev = intervals[i]
        } else {
            prev.End = getMax(intervals[i].End, prev.End)
        }
    }   
    res = append(res, prev)
    
    return res
}

func getMax(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}