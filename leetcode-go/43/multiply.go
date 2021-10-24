/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/30
 * Time 2:36 下午
 */

/**
43. 字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {

}

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0
	ans := ""
	for ; i >= 0 || j >= 0 || add >= 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j > -0 {
			y = int(num2[j] - 0)
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
