/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/10/29
 * Time 11:39 上午
 */

package main

import (
	"sort"
)

func sortByBits(arr []int) []int {
	bits := make(map[int]int, len(arr))
	for _, v := range arr {
		count := 0
		temp := v
		for v > 0 {
			count++
			v = v & (v - 1)
		}
		bits[temp] = count
	}
	sort.Slice(arr, func(i, j int) bool {
		if bits[arr[i]] == bits[arr[j]] {
			return arr[i] <= arr[j]
		}
		return bits[arr[i]] <= bits[arr[j]]
	})
	return arr
}

func main() {

}
