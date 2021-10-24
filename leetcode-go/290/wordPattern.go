/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/16
 * Time 5:54 下午
 */

/**
290. 单词规律
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
*/

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	//使用map
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}
	m := make(map[rune]string)
	mS := make(map[string]int)
	for i, v := range pattern {
		capital, ok := m[v]
		if !ok {
			//a a a a的情况
			// m中包含arr[i]
			if _, y := mS[arr[i]]; !y {
				mS[arr[i]] = 1
			} else {
				return false
			}
			m[v] = arr[i]
		} else {
			if capital != arr[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	pattern := "abba"
	str := "dog dog dog dog"
	fmt.Println(wordPattern(pattern, str))
}
