package skiena_2

import (
	"fmt"
	"sort"
	"testing"
)

/**
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]


Constraints:

1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/

func fourSum(nums []int, target int) [][]int {
	r := [][]int{}
	rxs := map[string]struct{}{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			} else {
				for k := 0; k < len(nums); k++ {
					if k == j {
						continue
					} else {
						for l := 0; l < len(nums); l++ {
							if l == k {
								continue
							} else {
								if nums[i]+nums[j]+nums[k]+nums[l] == target {
									xs := []int{nums[i], nums[j], nums[k], nums[l]}
									sort.Ints(xs)
									sxs := fmt.Sprint(xs)
									_, ok := rxs[sxs]
									if !ok {
										rxs[sxs] = struct{}{}
										r = append(r, xs)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return r
}
func TestFourSum(t *testing.T) {

	xs := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	//expected := [][]int{{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}}
	fmt.Println(xs)
	//if !arrays.ArrayEquality(xs, expected, eq) {
	//	t.Errorf("Actual:%v, Expected:%v", xs, expected)
	//}

	xs = fourSum([]int{2, 2, 2, 2, 2}, 8)
	//expected := [][]int{{2,2,2,2}}
	fmt.Println(xs)
	//if !arrays.ArrayEquality(xs, expected, eq) {
	//	t.Errorf("Actual:%v, Expected:%v", xs, expected)
	//}

}
