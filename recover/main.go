/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/26
 * Time 上午11:31
 */

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误", err)
		panic(err)
	}

	fmt.Println("上面的出发操作执行成功。。。")
	fmt.Println("正常执行下面的逻辑。。。")

}

func test() (err error) {
	//// 利用defer+recover来捕获错误：
	//defer func() {
	//	// 调用recover内置函数，可以捕获错误：
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("错误已经捕获")
	//		fmt.Println("err是：", err)
	//	}
	//}()
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出自定义错误
		return errors.New("除数不能为0哦～")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
