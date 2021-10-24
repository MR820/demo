/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/10/13
 * Time 上午10:04
 */

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func test() error {
	if _, file, line, ok := runtime.Caller(1); ok {
		return &MyError{"Something happened", file, line}
	}
	return &MyError{"Something happened", "", 0}

}

func main() {

	err := test()
	fmt.Printf("%T \n", errors.Cause(err))
	fmt.Printf("%v \n", errors.Cause(err))
	fmt.Printf("%T \n", errors.Wrap(err, "write failed"))
	fmt.Printf("%v \n", errors.Wrap(err, "write failed"))
	switch err := err.(type) {
	case nil:
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:

	}
}
