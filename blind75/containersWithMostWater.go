package blind75

// when calculating something between values, think two pointers
// first step of two pointers is almost always l, r := 0, len(<slice>)-1
// and the loop is then for l < r
// then some stage you have to do a l++, r--
// after that you just have to think about under what conditions l moves
// and r moves

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea, area := 0, 0
	for l < r {
		area = (r - l) * min(height[l], height[r])
		maxArea = max(area, maxArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}
