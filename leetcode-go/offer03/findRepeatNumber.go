/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/10
 * Time 11:33 上午
 */

package main

import "fmt"

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = i
	}
	return -1
}

func main() {
	s := []int{1, 2, 4, 3, 45, 3, 4, 72, 3}
	fmt.Println(findRepeatNumber(s))
}
