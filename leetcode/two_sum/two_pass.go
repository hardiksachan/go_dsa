package two_sum

func TwoPass(nums []int, target int) []int {
	lookup := make(map[int]int, 0)

	for i, v := range nums {
		lookup[v] = i
	}

	for i, v := range nums {
		j, ok := lookup[target-v]
		if i == j {
			continue
		}
		if ok {
			return []int{i, j}
		}
	}

	return []int{}
}
