package t13XX

import "sort"

//[idi, ratingi, veganFriendlyi, pricei, distancei]
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	ids := []int{}

	sort.Slice(restaurants,
		func(i, j int) bool {
			if restaurants[i][1] == restaurants[j][1] {
				return restaurants[i][0] > restaurants[j][0]
			}
			return restaurants[i][1] > restaurants[j][1]
		})

	for _, res := range restaurants {
		if (veganFriendly == 1 && res[2] != veganFriendly) || res[3] > maxPrice || res[4] > maxDistance {
			continue
		}
		ids = append(ids, res[0])
	}

	return ids
}
