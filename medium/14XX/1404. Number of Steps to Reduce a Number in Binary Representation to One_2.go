package _14XX

/*
1404. Number of Steps to Reduce a Number in Binary Representation to One
Medium

94

13

Add to List

Share
Given a number s in their binary representation. Return the number of steps to reduce it to 1 under the following rules:

If the current number is even, you have to divide it by 2.

If the current number is odd, you have to add 1 to it.

It's guaranteed that you can always reach to one for all testcases.



Example 1:

Input: s = "1101"
Output: 6
Explanation: "1101" corressponds to number 13 in their decimal representation.
Step 1) 13 is odd, add 1 and obtain 14.
Step 2) 14 is even, divide by 2 and obtain 7.
Step 3) 7 is odd, add 1 and obtain 8.
Step 4) 8 is even, divide by 2 and obtain 4.
Step 5) 4 is even, divide by 2 and obtain 2.
Step 6) 2 is even, divide by 2 and obtain 1.
Example 2:

Input: s = "10"
Output: 1
Explanation: "10" corressponds to number 2 in their decimal representation.
Step 1) 2 is even, divide by 2 and obtain 1.
Example 3:

Input: s = "1"
Output: 0
*/
func NumSteps_2(s string) int {
	steps := 0

	for s != "1" && len(s) != 0 {
		if s[len(s)-1] == '0' {
			s = s[0 : len(s)-1]
		} else {
			for i := len(s) - 1; i >= 0; i-- {
				if s[i] == '0' {
					suffix := s[i+1:]
					if i == len(s)-1 {
						suffix = ""
					}
					s = s[:i] + "1" + suffix
					break
				} else {
					suffix := s[i+1:]
					if i == len(s)-1 {
						suffix = ""
					}
					s = s[:i] + "0" + suffix
				}

			}
		}
		steps++
	}

	return steps
}