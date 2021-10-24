/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/15
 * Time 10:47 上午
 */

package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	length := len(s)
	if n <= 0 || n > length {
		return s
	}
	s1 := fmt.Sprintf("%s%s", string([]rune(s)[n:]), string([]rune(s)[:n]))
	return s1
}

func main() {
	s := "adew"
	fmt.Println(reverseLeftWords(s, 2))
}
