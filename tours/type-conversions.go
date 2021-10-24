/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/8/25
 * Time 3:59 下午
 */

package main

import (
	"fmt"
	"io"
	"strings"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		z := x + y
		x = y
		y = z
		return z
	}
}

func main() {
	//f := fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(f())
	//}
	//
	//fmt.Println(-math.Sqrt2)
	r := strings.NewReader("Hello World!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
