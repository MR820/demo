/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/13
 * Time 下午3:24
 */

package main

import "fmt"

/**
m中取两个 = m * (m-1)
简化为找到所有的m即可
*/

func main() {
	p := make([][]int, 0)
	p = append(p, []int{0, 0})
	p = append(p, []int{1, 0})
	p = append(p, []int{-1, 0})
	p = append(p, []int{0, 1})
	p = append(p, []int{0, -1})
	fmt.Println(numberOfBoomerangs(p))
}
func numberOfBoomerangs(points [][]int) (ans int) {
	num := 0
	for _, p := range points {
		m := map[int]int{}
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			m[dis]++
		}
		for _, v := range m {
			n := v * (v - 1)
			num += n
		}
	}
	return num
}
