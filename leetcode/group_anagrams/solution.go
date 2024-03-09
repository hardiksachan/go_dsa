package group_anagrams

func HashMap(strs []string) [][]string {
	lookup := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, c := range str {
			count[c-'a']++
		}
		lookup[count] = append(lookup[count], str)
	}

	result := make([][]string, 0)

	for _, v := range lookup {
		result = append(result, v)
	}

	return result
}
