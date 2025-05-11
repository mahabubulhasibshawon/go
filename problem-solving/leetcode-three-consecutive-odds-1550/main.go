package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
        // Loop through the array up to the third-to-last element
        for i := 0; i < len(arr) - 2; i++ {
            // Check if the current element and the next two elements are all
            // odd
            if (arr[i] % 2 == 1 && arr[i + 1] % 2 == 1 && arr[i + 2] % 2 == 1) {
                return true
            }
        }
        return false;
}

func main() {
    nums := []int{2, 3, 5, 7, 10}
    fmt.Println(threeConsecutiveOdds(nums)) // Output: true
}
