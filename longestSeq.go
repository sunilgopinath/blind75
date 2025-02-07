func longestConsecutive(nums []int) int {
	// make a map of each value

	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	// declare maxSeq variable
	maxSeq := 0
	// loop through, if the number starts a sequence, increment
	// else skip
	// compare curr Seq with max Seq
	for _, v := range nums {
		if m[v-1] {
			continue
		}
		currSeq, tmp := 1, v+1
		for m[tmp] {
			currSeq++
			tmp++
		}
		if currSeq > maxSeq {
			maxSeq = currSeq
		}
		if maxSeq > len(nums)/2 {
			break
		}
	}
	return maxSeq
}