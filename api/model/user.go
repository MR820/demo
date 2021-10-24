/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/8/28
 * Time 1:37 下午
 */

package model

type User struct {
	Id   int    `json:"id" uri:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
