/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/10/13
 * Time 上午9:50
 */

package main

import (
	"errors"
	"fmt"
)

func main() {
	Check(0)
}

func Check(n int) {
	pos, err := Positive(n)
	if err != nil {
		//fmt.Println(reflect.TypeOf(err.Error()))
		//fmt.Println(err.Error())
		fmt.Println(n, err)
		return
	}
	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func Positive(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}
	return n > -1, nil
}
