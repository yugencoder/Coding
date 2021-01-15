package t13XX

import (
	"fmt"
	"sort"
)

/*
A cinema has n rows of seats, numbered from 1 to n and there are ten seats in each row, labelled from 1 to 10 as shown in the figure above.

Given the array reservedSeats containing the numbers of seats already reserved, for example, reservedSeats[i]=[3,8] means the seat located in row 3 and labelled with 8 is already reserved.

Return the maximum number of four-person families you can allocate on the cinema seats. A four-person family occupies fours seats in one row, that are next to each other. Seats across an aisle (such as [3,3] and [3,4]) are not considered to be next to each other, however, It is permissible for the four-person family to be separated by an aisle, but in that case, exactly two people have to sit on each side of the aisle.



Example 1:



Input: n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
Output: 4
Explanation: The figure above shows the optimal allocation for four families, where seats mark with blue are already reserved and contiguous seats mark with orange are for one family.
*/

func MaxNumberOfFamilies(x int, reservedSeats [][]int) int {

	g := 0b1
	fmt.Printf("%v", int(g))
	fmt.Printf("%v", g)

	m := len(reservedSeats)
	res := x * 2

	xLoc := map[string]bool{}

	sort.Slice(reservedSeats, func(i, j int) bool {
		return reservedSeats[i][0] < reservedSeats[j][0]
	})

	prev := reservedSeats[0][0]

	for i := 0; i <= m; i++ {

		if i == m || reservedSeats[i][0] != prev {
			if xLoc["lm"] && xLoc["rm"] ||
				xLoc["lm"] && xLoc["r"] ||
				xLoc["l"] && xLoc["rm"] {
				res -= 2

			} else {
				if (xLoc["l"] || xLoc["r"]) && !xLoc["lm"] && !xLoc["rm"] {
					res -= 1
				} else {
					if (xLoc["l"] || xLoc["lm"]) && !xLoc["rm"] && !xLoc["r"] {
						res -= 1
					}
					if (xLoc["r"] || xLoc["rm"]) && !xLoc["lm"] && !xLoc["l"] {
						res -= 1
					}
				}
			}

			if i == m {
				break
			}
			xLoc = map[string]bool{}
			prev = reservedSeats[i][0]
		}

		col := reservedSeats[i][1]

		if col > 1 && col < 4 {
			xLoc["l"] = true
		}

		if col > 7 && col < 10 {
			xLoc["r"] = true
		}

		if col > 3 && col < 6 {
			xLoc["lm"] = true
		}

		if col > 5 && col < 8 {
			xLoc["rm"] = true
		}

	}

	return res
}
