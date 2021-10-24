/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/28
 * Time 5:17 下午
 */

/**
78. 子集
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

package main

import "fmt"

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < 1<<n; i++ {
		set := []int{}
		for j, v := range nums {
			if i>>j&1 > 0 {
				set = append(set, v)
			}

		}
		ans = append(ans, append([]int(nil), set...))
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
