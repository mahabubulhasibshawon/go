package main

import "fmt"

const MOD = 1000000007

func modPow(base, exp, mod int) int {
	res := 1
	base = base % mod

	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp = exp / 2
	}
	return res
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, m, d int
		fmt.Scan(&n, &m, &d)

		nonFixed := n - m
		ans := modPow(d, nonFixed, MOD)
		fmt.Println(ans)

	}
}

/*
While exploring the Best Golang Community Ever (BGCE) Discord server, you found a mysterious notebook describing alien
 number systems. For example, Earth’s decimal system has digits 0–9, but other planets might use systems with 2, 3, or 
 even up to 16 distinct symbols. Recently, Bangladesh’s mobile operators (like Grameenphone) introduced new number series
  prefixes (such as 013) to expand their pool of SIM numbers. But why add new prefixes? Because:

Example: the 017 series →

Total length = 11

First 3 digits fixed (017)

Remaining 8 digits → only 10^8 combinations (decimal system)

You have a wild idea: What if we switch to a different number system (from the alien notebook)? Could we avoid adding 
new series by dramatically increasing the combinations?

Task:

N → total length of the phone number
M → number of fixed digits at the front (prefix)
D → number of unique digits in the number system (base)
Your task is to calculate the total number of possible combinations using this number system. Since the result can be large, 
output it modulo 10^9+7.

Input Format

Each test contains multiple test cases. The first line contains the number of test cases T (1<=t<=1350) The description 
of the test cases follows. For each test case contains three integers given N,M,D (1<=N<=20,1<=M<=min(5,N),2<=D<=16) the 
length of phone number, total fixed digits at front and number system base (number of distinct digits)

Constraints

T (1<=t<=1350) N,M,D (1<=N<=20,1<=M<=min(5,N),2<=D<=16)

Output Format

For each test case, print one line: the total possible combinations modulo 10^9+7.

Sample Input 0

5
1 1 2
2 1 2
11 3 10
11 2 10
11 5 9
Sample Output 0

1
2
100000000
1000000000
531441
Explanation 0

Note: In the first test case, the length and fixed digits are the same, so the only possible combination is 1.
 In the second test case, the first digit is fixed, so in the second digit we can enter either 1 or 0; the answer is 2. 
 The third test case is described in the description.
*/
