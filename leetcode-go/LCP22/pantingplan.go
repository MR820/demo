/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/23
 * Time 5:37 下午
 */

/**
LCP 22. 黑白方格画
小扣注意到秋日市集上有一个创作黑白方格画的摊位。摊主给每个顾客提供一个固定在墙上的白色画板，画板不能转动。画板上有 n * n 的网格。绘画规则为，小扣可以选择任意多行以及任意多列的格子涂成黑色，所选行数、列数均可为 0。

小扣希望最终的成品上需要有 k 个黑色格子，请返回小扣共有多少种涂色方案。

注意：两个方案中任意一个相同位置的格子颜色不同，就视为不同的方案。

示例 1：

输入：n = 2, k = 2

输出：4

解释：一共有四种不同的方案：
第一种方案：涂第一列；
第二种方案：涂第二列；
第三种方案：涂第一行；
第四种方案：涂第二行。

示例 2：

输入：n = 2, k = 1

输出：0

解释：不可行，因为第一次涂色至少会涂两个黑格。

示例 3：

输入：n = 2, k = 4

输出：1

解释：共有 2*2=4 个格子，仅有一种涂色方案。

限制：

1 <= n <= 6
0 <= k <= n * n
*/

package main

import "fmt"

func paintingPlan(n int, k int) int {
	if k == 0 || k == n*n {
		return 1
	}
	if k < n {
		return 0
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := i*n + j*n - i*j
			if sum == k {
				x := factorial(n, i)
				y := factorial(n, j)
				ans += x * y
			}
		}
	}
	return ans

}

func factorial(n int, m int) int {
	ans := 1
	for i := n; i > m; i-- {
		ans *= i
	}
	for i := n - m; i > 0; i-- {
		ans /= i
	}
	return ans
}

func main() {
	n := 2
	k := 2
	fmt.Println(paintingPlan(n, k))
}
