package group_anagrams_test

import (
	"testing"

	"github.com/hardik-sachan/go_dsa/leetcode/group_anagrams"
	"github.com/hardik-sachan/go_dsa/utils"
)

type testcase struct {
	strs   []string
	target [][]string
}

var testCases = []testcase{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	{[]string{""}, [][]string{{""}}},
	{[]string{"a"}, [][]string{{"a"}}},
}

func TestHashMap(t *testing.T) {
	for _, tc := range testCases {
		got := group_anagrams.HashMap(tc.strs)
		if !utils.SliceOfSliceEqualWithoutOrder(got, tc.target) {
			t.Errorf("Expected %v, got %v", tc.target, got)
		}
	}
}
