/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/2/8
 * Time 3:18 下午
 */

/**
978. 最长湍流子数组
当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 A 的最大湍流子数组的长度。



示例 1：

输入：[9,4,2,10,7,8,8,1,9]
输出：5
解释：(A[1] > A[2] < A[3] > A[4] < A[5])
示例 2：

输入：[4,8,12,16]
输出：2
示例 3：

输入：[100]
输出：1


提示：

1 <= A.length <= 40000
0 <= A[i] <= 10^9
*/

package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	start, maxLen, end := 0, 0, 0
	l := len(arr)
	if l == 0 || l == 1 {
		return l
	}
	//true 规则1 false 规则2
	flag := true
	for end < l-1 {
		if (end%2 == 0 && flag == true && arr[end] >= arr[end+1]) ||
			(end%2 == 0 && flag == false && arr[end] <= arr[end+1]) ||
			(end%2 != 0 && flag == true && arr[end] <= arr[end+1]) ||
			(end%2 != 0 && flag == false && arr[end] >= arr[end+1]) {
			start = end
			flag = reverse(flag)
		}
		end++
		if arr[start] == arr[end] {
			maxLen = max(maxLen, 1)
		} else {
			maxLen = max(maxLen, end-start+1)
		}

	}
	return maxLen
}

func reverse(b bool) bool {
	if b == true {
		return false
	} else {
		return true
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := []int{9, 9}
	fmt.Println(maxTurbulenceSize(s))
}
