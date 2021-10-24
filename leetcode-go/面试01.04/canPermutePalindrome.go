/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/2
 * Time 4:34 下午
 */

/**
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-permutation-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func canPermutePalindrome(s string) bool {
	//回文排列的条件是最多只有一个字符为奇数个

	m := make(map[int32]int)

	for _, v := range s {
		//字符串遍历变为ascii码
		//fmt.Println(v)
		//fmt.Println(reflect.TypeOf(v))
		capital, ok := m[v]
		if ok {
			m[v] = capital + 1
		} else {
			m[v] = 1
		}
	}

	n := 0
	// 遍历m，确定奇数的个数
	for cap := range m {
		if m[cap]%2 != 0 {
			n++
		}
	}
	if n > 1 {
		return false
	} else {
		return true
	}
}

func main() {
	s := "cooc"
	fmt.Print(canPermutePalindrome(s))
}
