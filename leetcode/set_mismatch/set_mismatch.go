package setmismatch

func FindErrorNums(nums []int) []int {
	n := len(nums)
	counts := make([]int, n+1)

	for i := 0; i < n; i++ {
		counts[nums[i]]++
	}

	ret := []int{0, 0}

	for j := 1; j < n+1; j++ {
		if counts[j] > 1 {
			ret[0] = j
		}
		if counts[j] == 0 {
			ret[1] = j
		}
	}
	return ret
}
