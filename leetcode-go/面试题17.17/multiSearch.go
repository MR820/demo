/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/13
 * Time 2:00 下午
 */

package main

import (
	"fmt"
)

/**
面试题 17.17. 多次搜索
给定一个较长字符串big和一个包含较短字符串的数组smalls，设计一个方法，根据smalls中的每一个较短字符串，对big进行搜索。输出smalls中的字符串在big里出现的所有位置positions，其中positions[i]为smalls[i]出现的所有位置。

示例：

输入：
big = "mississippi"
smalls = ["is","ppi","hi","sis","i","ssippi"]
输出： [[1,4],[8],[],[3],[1,4,7,10],[5]]
提示：

0 <= len(big) <= 1000
0 <= len(smalls[i]) <= 1000
smalls的总字符数不会超过 100000。
你可以认为smalls中没有重复字符串。
所有出现的字符均为英文小写字母。
*/

func multiSearch(big string, smalls []string) [][]int {
	// 前缀树的理想状态是左前缀确定，pass标记路过几次，end标记以此为结尾
	// 此处对pass和end都没有使用
	//直接暴力破解
	length := len(big)
	arr := make([][]int, 0)
	for _, v := range smalls {
		l := len(v)
		s := make([]int, 0)
		if l == 0 {
			arr = append(arr, s)
			continue
		}
		for i := 0; i <= length-l; i++ {
			if big[i:i+l] == v {
				s = append(s, i)
			}
		}
		arr = append(arr, s)
	}
	return arr
}

func main() {
	big := "mississippi"
	smalls := []string{"is", "ppi", "hi", "sis", "i", "ssippi"}
	fmt.Println(multiSearch(big, smalls))
}
