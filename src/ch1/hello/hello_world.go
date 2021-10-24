/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/10/12
 * Time 6:20 下午
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World!", os.Args[1])
	}

}
