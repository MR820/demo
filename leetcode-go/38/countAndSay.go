/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/11
 * Time 11:09 上午
 */

package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	s := "11"
	for j := 3; j <= n; j++ {
		var r string
		count := 1
		for i := 1; i <= len(s)-1; i++ {
			if s[i] != s[i-1] {
				r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i-1]))
				count = 1
			} else {
				count++
			}
			if i+1 == len(s) {
				r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i]))
				s = r
				break
			}
		}
	}

	return s
}

func main() {
	fmt.Println(countAndSay(9))
}
