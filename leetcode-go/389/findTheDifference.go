/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/18
 * Time 3:57 下午
 */

/**
389. 找不同
给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。



示例 1：

输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。
示例 2：

输入：s = "", t = "y"
输出："y"
示例 3：

输入：s = "a", t = "aa"
输出："a"
示例 4：

输入：s = "ae", t = "aea"
输出："a"


提示：

0 <= s.length <= 1000
t.length == s.length + 1
s 和 t 只包含小写字母
*/

package main

import "fmt"

func findTheDifference(s string, t string) byte {
	arr := [26]int{}
	var res rune
	for _, v := range s {
		arr[v-'a']++
	}
	for _, v := range t {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			res = v
		}
	}
	return byte(res)
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(findTheDifference(s, t))
	fmt.Println(string(findTheDifference(s, t)))
}
