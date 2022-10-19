func twoSum(nums []int, target int) []int {
	Ans := []int{0, 1}
	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			if a+nums[j] == target {
				Ans[0] = i
				Ans[1] = j
				break
			}
		}
	}
	return Ans
}