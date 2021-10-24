/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/14
 * Time 4:46 下午
 */

package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	s := make([]bool, 0)
	for i := range candies {
		if candies[i]+extraCandies >= max {
			s = append(s, true)
		} else {
			s = append(s, false)
		}
	}
	return s
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
