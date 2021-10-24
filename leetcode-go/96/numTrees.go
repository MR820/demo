/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/1/19
 * Time 1:29 下午
 */

/**
96. 不同的二叉搜索树
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

package main

import "fmt"

/**
二叉搜索树条件
left < root < right
*/
func numTrees(n int) int {
	//动态规划
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

func main() {
	fmt.Println(numTrees(3))
}
