/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/6
 * Time 下午6:28
 */

package group

import "fmt"

type Group struct {
	Id   int
	Name string
	Pid  int
}

func (group Group) Test() {
	fmt.Printf("group name %v,id %v ,pid %v", group.Name, group.Id, group.Pid)
}
