/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/2
 * Time 5:03 下午
 */

/**
给定一个增序排列数组 nums ，你需要在 原地 删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {
	// 空间复杂度为O(1)
	length := len(nums)

	i, j := 1, 2
	for ; j < length; j++ {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1

}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 4}
	fmt.Println(removeDuplicates(nums))
}
