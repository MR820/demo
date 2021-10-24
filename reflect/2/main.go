/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/6
 * Time 下午6:12
 */

package main

import (
	"imooc.com/ccmouse/learngo/reflect/2/group"
	"imooc.com/ccmouse/learngo/reflect/2/server"
	"imooc.com/ccmouse/learngo/reflect/2/user"
)

func main() {
	reServer := &server.ReServer{
		M: make(map[string]interface{}),
	}
	user := &user.User{
		Id:   1,
		Name: "XZW",
		Age:  20,
	}
	reServer.RegisterService(user)

	group := &group.Group{
		Id:   121,
		Name: "developer",
		Pid:  0,
	}
	reServer.RegisterService(group)

	reServer.Start("Group", "Test")
}
