
# 1550. Three Consecutive Odds

## Problem Statement

Given an integer array `arr`, return `true` if there are **three consecutive odd numbers** in the array. Otherwise, return `false`.

### Example 1:

```
Input: arr = [2, 6, 4, 1]  
Output: false  
Explanation: There are no three consecutive odds.
```

### Example 2:

```
Input: arr = [1, 2, 34, 3, 4, 5, 7, 23, 12]  
Output: true  
Explanation: [5, 7, 23] are three consecutive odd numbers.
```

## Constraints

* `1 <= arr.length <= 1000`
* `1 <= arr[i] <= 1000`

## Approach

* Traverse the array while checking for three consecutive odd numbers.
* Use a counter to keep track of consecutive odd numbers.
* If an even number is encountered, reset the counter.
* Return `true` if the counter reaches 3; otherwise, return `false`.


