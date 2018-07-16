package medianOfTwoSortedArrays

// There are two sorted arrays nums1 and nums2 of size m and n respectively.

// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

// Example 1:
// nums1 = [1, 3]
// nums2 = [2]

// The median is 2.0

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    totalNumber := len(nums1) + len(nums2)
    if totalNumber%2 == 0 {
        return float64(findKthNumber(nums1, nums2, totalNumber/2) + findKthNumber(nums1, nums2, totalNumber/2+1))/2
    }
    return float64(findKthNumber(nums1, nums2, totalNumber/2+1))
}

func findKthNumber(nums1 []int, nums2 []int, k int) int {
    if len(nums1) == 0 { return nums2[k-1] }
    if len(nums2) == 0 { return nums1[k-1] }
    if k == 1 {return findMin(nums1[0], nums2[0])}
    
    if len(nums1) < k/2 { return findKthNumber(nums1, nums2[k/2:], k - k/2) }
    if len(nums2) < k/2 { return findKthNumber(nums1[k/2:], nums2, k - k/2) }
    
    if nums1[k/2-1] > nums2[k/2-1] {
        return findKthNumber(nums1, nums2[k/2:], k - k/2)
    } else {
        return findKthNumber(nums1[k/2:], nums2, k - k/2)
    }
}

func findMin(a int, b int) int {
    if a > b { return b}
    return a
}