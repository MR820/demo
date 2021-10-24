/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/31
 * Time 3:47 下午
 */

/**
204. 计数质数
统计所有小于非负整数 n 的质数的数量。



示例 1：

输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
示例 2：

输入：n = 0
输出：0
示例 3：

输入：n = 1
输出：0


提示：

0 <= n <= 5 * 106
*/

package main

import "fmt"

/**
func countPrimes(n int) int {
	//暴力破解超时
	count := 0
	for i := 2; i < n; i++ {
		if isPrimes(i) {
			count++
		}
	}
	return count
}

func isPrimes(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

*/

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	ret := 0
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] == true {
			ret++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(countPrimes(10))
}
