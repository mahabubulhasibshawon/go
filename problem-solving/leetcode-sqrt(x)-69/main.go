package main

import (
    "fmt"
)

func mySqrt(x int) int {
	// // brute force solution
	// // Time complexity: O(n)
	// var res int
    // for i := 1; i <= x/2; i++ {
	// 	if (i*i) <= x {
	// 		res = i
	// 	}
	// }
	// return res
	//     EDGE CASES
    if x == 0 || x == 1 {
        return x
    }
    
    low :=1
    high :=x
    
    for low <= high {
        mid := (low + high) / 2
        
        if (mid * mid) == x {
            return mid
        } else if (mid * mid) < x {
            low = mid + 1
        } else {
            high = mid -1
        }
    }
	return low - 1
}

func main() {
    inputs := []int{0, 1, 4, 8, 10, 15, 100, 1000}
    for _, x := range inputs {
        fmt.Printf("mySqrt(%d) = %d\n", x, mySqrt(x))
    }
}
