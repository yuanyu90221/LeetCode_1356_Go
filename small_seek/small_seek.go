package small_seek

func SmallerNumnersThanCurrent(nums []int) []int {
	result := []int{}
	for idx := range nums {
		result = append(result, countSmall(nums, idx))
	}
	return result
}

func countSmall(nums []int, idx int) int {
	result := 0
	for i, val := range nums {
		if i == idx {
			continue
		}
		if val < nums[idx] {
			result += 1
		}
	}
	return result
}
