func twoSum(nums []int, target int) []int {
	var indices []int
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				indices = append(indices, i, j)
				return indices
			}
		}
	}
	return nil
}
