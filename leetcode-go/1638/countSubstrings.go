/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/10
 * Time 2:56 下午
 */

/**
1638. 统计只差一个字符的子串数目
给你两个字符串 s 和 t ，请你找出 s 中的非空子串的数目，这些子串满足替换 一个不同字符 以后，是 t 串的子串。换言之，请你找到 s 和 t 串中 恰好 只有一个字符不同的子字符串对的数目。

比方说， "computer" 和 "computation" 加粗部分只有一个字符不同： 'e'/'a' ，所以这一对子字符串会给答案加 1 。

请你返回满足上述条件的不同子字符串对数目。

一个 子字符串 是一个字符串中连续的字符。



示例 1：

输入：s = "aba", t = "baba"
输出：6
解释：以下为只相差 1 个字符的 s 和 t 串的子字符串对：
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
加粗部分分别表示 s 和 t 串选出来的子字符串。
示例 2：
输入：s = "ab", t = "bb"
输出：3
解释：以下为只相差 1 个字符的 s 和 t 串的子字符串对：
("ab", "bb")
("ab", "bb")
("ab", "bb")
加粗部分分别表示 s 和 t 串选出来的子字符串。
示例 3：
输入：s = "a", t = "a"
输出：0
示例 4：

输入：s = "abe", t = "bbc"
输出：10


提示：

1 <= s.length, t.length <= 100
s 和 t 都只包含小写英文字母。

*/

package main

import "fmt"

func countSubstrings(s string, t string) int {
	//暴力破解
	l := 0
	if len(s) < len(t) {
		l = len(s)
	} else {
		l = len(t)
	}
	//子字符串从1～l
	//i记录s的开始位置，j记录t的开始位置
	count := 0
	//长度为k
	for k := 1; k <= l; k++ {
		for i := 0; i <= len(s)-k; i++ {
			for j := 0; j <= len(t)-k; j++ {
				if compare(s[i:i+k], t[j:j+k]) {
					count++
				}
			}
		}
	}
	return count
}

func compare(s string, t string) bool {
	//s和t长度确保相等
	l := len(s)
	diff := 0
	for i := 0; i < l; i++ {
		if s[i] != t[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	if diff == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "aba"
	t := "baba"
	fmt.Println(countSubstrings(s, t))
}
