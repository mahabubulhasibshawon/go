// Leetcode Problem 704: Binary Search
// https://leetcode.com/problems/binary-search/description/

package main

import "fmt"

func binarySearch(arr []int, target int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := low + (high-low)/2

        if arr[mid] == target {
            return mid // target found
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1 // target not found
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13}
    target := 7

    result := binarySearch(arr, target)

    if result != -1 {
        fmt.Printf("Element found at index %d\n", result)
    } else {
        fmt.Println("Element not found")
    }
}

/*
704. Binary Search
Solved
Easy
Topics
Companies
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
 

Constraints:

1 <= nums.length <= 104
-104 < nums[i], target < 104
All the integers in nums are unique.
nums is sorted in ascending order.
*/