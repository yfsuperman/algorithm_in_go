package maximumBinaryTree

// Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

// The root is the maximum number in the array.
// The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
// The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
// Construct the maximum tree by the given array and output the root node of this tree.

// Example 1:
// Input: [3,2,1,6,0,5]
// Output: return the tree root node representing the following tree:

//       6
//     /   \
//    3     5
//     \    / 
//      2  0   
//        \
//         1


type TreeNode struct {
	Val int
    Left *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    
    if len(nums) == 0 {
        return nil
    }
    
    index, maxNum := 0, 0
    for i, num := range nums {
        if num > maxNum {
            maxNum = num
            index = i
        }
    }
    
	return &TreeNode{maxNum, 
		constructMaximumBinaryTree(nums[:index]), 
		constructMaximumBinaryTree(nums[index+1:])}
    
}