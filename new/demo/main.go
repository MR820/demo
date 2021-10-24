/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/10
 * Time 上午9:32
 */

/**
new 和 make 的区别
nil slice 和 空 slice 的区别
*/

package main

import (
	"fmt"
)

func main() {
	var v int // 声明变量
	fmt.Println(v)
	v = 8
	fmt.Println(v)

	var val *int // 声明指针变量,值为nil
	fmt.Println(val)
	val = new(int)
	*val = 8
	fmt.Println(val)
	fmt.Println(*val)

	var slice []int
	fmt.Println(slice)

}
