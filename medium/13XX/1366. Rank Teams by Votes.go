package t13XX

import (
	"fmt"
	"sort"
	"strconv"
)

/*

In a special ranking system, each voter gives a rank from highest to lowest to all teams participated in the competition.

The ordering of teams is decided by who received the most position-one votes. If two or more teams tie in the first position, we consider the second position to resolve the conflict, if they tie again, we continue this process until the ties are resolved. If two or more teams are still tied after considering all positions, we rank them alphabetically based on their team letter.

Given an array of strings votes which is the votes of all voters in the ranking systems. Sort all teams according to the ranking system described above.

Return a string of all teams sorted by the ranking system.



Example 1:

Input: votes = ["ABC","ACB","ABC","ACB","ACB"]
Output: "ACB"
Explanation: Team A was ranked first place by 5 voters. No other team was voted as first place so team A is the first team.
Team B was ranked second by 2 voters and was ranked third by 3 voters.
Team C was ranked second by 3 voters and was ranked third by 2 voters.
As most of the voters ranked C second, team C is the second team and team B is the third.
Example 2:

Input: votes = ["WXYZ","XYZW"]
Output: "XWYZ"
Explanation: X is the winner due to tie-breaking rule. X has same votes as W for the first position but X has one vote as second position while W doesn't have any votes as second position.
Example 3:

Input: votes = ["ZMNAGUEDSJYLBOPHRQICWFXTVK"]
Output: "ZMNAGUEDSJYLBOPHRQICWFXTVK"
Explanation: Only one voter so his votes are used for the ranking.
Example 4:

Input: votes = ["BCA","CAB","CBA","ABC","ACB","BAC"]
Output: "ABC"
Explanation:
Team A was ranked first by 2 voters, second by 2 voters and third by 2 voters.
Team B was ranked first by 2 voters, second by 2 voters and third by 2 voters.
Team C was ranked first by 2 voters, second by 2 voters and third by 2 voters.
There is a tie and we rank teams ascending by their IDs.
Example 5:

Input: votes = ["M","M","M","M"]
Output: "M"
Explanation: Only team M in the competition so it has the first rank.
*/

type rankData struct {
	sequence string
	key      string
}

func RankTeams(votes []string) string {
	var result string
	rankMap := map[string]*rankData{}
	//for each of the letters - have n score.

	for _, vote := range votes {
		for i, s := range vote {
			s := string(s)
			if _, ok := rankMap[s]; !ok {
				rankMap[s] = &rankData{}
				rankMap[s].key = s
				rankMap[s].sequence = "                                                                              "
			}

			val, _ := strconv.Atoi(rankMap[s].sequence[3*i : 3*i+3])
			val++
			rankMap[s].sequence = rankMap[s].sequence[:3*i] + fmt.Sprintf("%03d", val) + rankMap[s].sequence[3*i+3:]
		}
	}

	ranks := collect(rankMap)
	sort.Sort(Ranks(ranks))

	for _, r := range ranks {
		result += r.key
	}
	return result
}

func collect(rankMap map[string]*rankData) []*rankData {
	var ranks []*rankData

	for _, value := range rankMap {
		ranks = append(ranks, value)
	}

	return ranks
}

type Ranks []*rankData

func (a Ranks) Len() int { return len(a) }

func (a Ranks) Less(i, j int) bool {
	if a[i].sequence == a[j].sequence {
		return a[i].key < a[j].key
	}
	return a[i].sequence > a[j].sequence
}

func (a Ranks) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
