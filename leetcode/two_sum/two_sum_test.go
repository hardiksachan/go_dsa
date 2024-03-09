package two_sum_test

import (
	"testing"

	"github.com/hardik-sachan/go_dsa/leetcode/two_sum"
	"github.com/hardik-sachan/go_dsa/utils"
)

type testcase struct {
	nums   []int
	target int
	result []int
}

var testCases = []testcase{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestBrute(t *testing.T) {
	for _, tc := range testCases {
		got := two_sum.Brute(tc.nums, tc.target)
		if !utils.SliceEqual(got, tc.result) {
			t.Errorf("Expected %v, got %v", tc.result, got)
		}
	}
}

func TestTwoPass(t *testing.T) {
	for _, tc := range testCases {
		got := two_sum.TwoPass(tc.nums, tc.target)
		if !utils.SliceEqual(got, tc.result) {
			t.Errorf("Expected %v, got %v", tc.result, got)
		}
	}
}

func TestOnePass(t *testing.T) {
	for _, tc := range testCases {
		got := two_sum.OnePass(tc.nums, tc.target)
		if !utils.SliceEqual(got, tc.result) {
			t.Errorf("Expected %v, got %v", tc.result, got)
		}
	}
}
