/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/27
 * Time 2:04 下午
 */

/**
394. 字符串解码
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。



示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stack struct {
	size int
	s    []string
}

func Constructor() *stack {
	return &stack{
		size: 0,
		s:    nil,
	}
}

func (this *stack) add(c string) {
	this.size++
	this.s = append(this.s, c)
}

func (this *stack) pop() string {
	if this.size <= 0 {
		panic("栈空")
	}
	this.size--
	tmp := this.s[this.size]
	this.s = this.s[:this.size]
	return tmp
}

func (this *stack) isEmpty() bool {
	if this.size == 0 {
		return true
	} else {
		return false
	}
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for s[*ptr] >= '0' && s[*ptr] <= '9' {
		ret += string(s[*ptr])
		*ptr++
	}
	return ret
}

func (this *stack) getString() string {
	s := ""
	for !this.isEmpty() {
		c := this.pop()
		if c == "[" {
			return s
		} else {
			s = c + s
		}
	}
	return s
}

func (this *stack) getRepTime() int {
	n := 0
	if !this.isEmpty() {
		s := this.pop()
		n, _ = strconv.Atoi(s)
	}
	return n
}

func decodeString(s string) string {
	stack := Constructor()
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stack.add(digits)
		} else if cur != ']' {
			stack.add(string(cur))
			ptr++
		} else {
			ptr++
			s := stack.getString()
			n := stack.getRepTime()
			t := strings.Repeat(s, n)
			stack.add(t)
		}
	}
	return stack.getString()
}

func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}
