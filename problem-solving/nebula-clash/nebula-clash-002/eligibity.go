package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var week [7]int
		valid := true
		attendance := 0

		for j := 0; j < 7; j++ {
			fmt.Scan(&week[j])
			if week[j] > 1 {
				valid = false
			}
			attendance += week[j]
		}

		if valid && attendance >= 5 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
Welcome to BGCE (Best Golang Community Ever)! Here, people can find the best resources to learn Golang, improve their 
problem-solving skills, prepare for coding interviews, and develop a deep understanding of operating systems and low-level 
programming.

Today, the big shot of the server, Kraken, has decided to take a mock interview of everyone. Naturally, everyone got super 
excited. But Kraken, being the deep-sea monster that he is, won’t just let anyone take the interview. He announced that only 
those who have joined our CP sessions at least 5 times last week, will be eligible.

Now here’s the twist: Kraken is terrible at CP. So, he asked his trusted sidekick, Starboy, to write a program to check who’s 
eligible. The problem? Starboy is kind of a noob at CP. And to make matters worse, all the CP pros on the server are mysteriously 
missing!

That’s where you come in — time to save the day.

You will be given an integer t (the number of test cases). For each test case, you’ll be given an array containing exactly 7 
integers, representing each day of the week. Each element will either be:

1 → the person attended a CP session that day

0 → the person missed the CP session

But watch out! Some sneaky people think they can outsmart the system and might try to enter numbers greater than 1. 
If any value in the array is greater than 1, that person is immediately disqualified and considered not eligible.

Your task is to determine whether each person is eligible for the interview or not.

Input Format

The first line contains an integer t — the number of test cases.

For each test case, a single line follows containing 7 space-separated integers:
a₁ a₂ a₃ a₄ a₅ a₆ a₇ where each aᵢ represents whether the person attended (1), missed (0), or possibly faked (any integer) a 
CP session on that day.

Constraints

1 ≤ t ≤ 2000 (as there are 2k members in the server)

For each test case, there are exactly 7 integers aᵢ where aᵢ ∈ ℤ (any integer, but ideally only 0 or 1 unless they're 
trying to be slick 😏)

Output Format

For each test case, output a single line:

"YES" (without quotes) if the person is eligible (i.e., attended 5 or more sessions and no value is greater than 1).

"NO" otherwise.

Sample Input 0

3
1 1 1 1 1 0 0
1 1 1 2 0 0 0
1 1 1 1 0 0 0
Sample Output 0

YES
NO
NO
Explanation 0

so for the first testcase, the person was present 5 times last week, so the output would be YES. for the second one, 
the person thought he is slick and put 2 for a day instead of 1. so he will not be eligible for the third case, the 
person only attended 4 times last week, so he is not eligible



*/