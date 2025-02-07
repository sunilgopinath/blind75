package blind75

// twoSum from leetcode
func twoSum(nums []int, target int) []int {
	// make a map
	m := make(map[int]int)
	// loop through looking for target - number in the map
	// if not, add into the map
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return []int{-1, -1}
}
