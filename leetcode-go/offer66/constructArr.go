/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/20
 * Time 4:54 下午
 */

/**
剑指 Offer 66. 构建乘积数组
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。



示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]


提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000
*/
package main

import "fmt"

/**
超出时限
*/
/**

func constructArr(a []int) []int {
	length := len(a)
	b := make([]int, 0)

	for i := 0; i < length; i++ {
		s := make([]int, 0)
		for j := 0; j < length; j++ {
			if i == j {
				continue
			}
			s = append(s, a[j])
		}

		product := 1
		for _, v := range s {
			product *= v
		}
		b = append(b, product)
	}
	return b
}

*/

//O(N^2)超出时限，想O(N),O(N*logN)
func constructArr(a []int) []int {
	//上下三角
	if len(a) == 0 {
		return []int{}
	}
	l := len(a)
	b := make([]int, l)
	b[0] = 1
	tmp := 1
	for i := 1; i < l; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	for i := l - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(a))
}
