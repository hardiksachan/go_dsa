package two_sum

func Brute(nums []int, target int) []int {
	for i, u := range nums {
		for j, v := range nums[i+1:] {
			if u+v == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{-1}
}
