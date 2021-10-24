/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/6
 * Time 下午6:16
 */

package user

import "fmt"

type User struct {
	Id   int
	Name string
	Age  int
}

func (user *User) FunOne() {
	fmt.Printf("user name %v,id %v,age %v", user.Name, user.Id, user.Age)
}
