# LeetCode 69: Sqrt(x)

## Problem Statement

Given a non-negative integer `x`, return the square root of `x` rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use `pow(x, 0.5)` in C++ or `x ** 0.5` in Python.

### Examples

#### Example 1:
**Input:**  
`x = 4`  
**Output:**  
`2`  
**Explanation:**  
The square root of 4 is 2, so we return 2.

#### Example 2:
**Input:**  
`x = 8`  
**Output:**  
`2`  
**Explanation:**  
The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

### Constraints

- `0 <= x <= 2³¹ - 1`

---

## Solution Approach

### Brute Force

Always start with the simplest, most naive approach to get a feel for the problem.

1. Try all numbers from `1` to `x`.
2. For each `i`, check if `i * i <= x`.
3. Stop when `i * i > x`.
4. Return the last valid `i`.

🔎 **Time Complexity:**  
This approach works, but if `x` is large (e.g., up to 10⁹), it could take up to 30,000 iterations. While not terrible, it is not optimal.

---

### Binary Search (Optimized)

Can we do better than `O(√x)`?  
Yes, we can exploit the fact that the square of a number increases rapidly. This makes it a great candidate for **Binary Search**, which is optimal for sorted, monotonic patterns.

#### Steps:
1. Search between `1` and `x`.
2. At each step, check if `mid * mid == x`.
3. If not, adjust the range accordingly.
4. Continue until the range converges.

🔎 **Time Complexity:**  
Binary Search reduces the number of iterations significantly compared to the brute force approach.

---