package two_sum

func OnePass(nums []int, target int) []int {
	lookup := make(map[int]int, 0)

	for i, v := range nums {
		j, ok := lookup[target-v]
		if ok {
			return []int{j, i}
		}
		lookup[v] = i
	}

	return []int{}
}
