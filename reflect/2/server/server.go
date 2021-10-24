/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/6
 * Time 下午6:02
 */

package server

import (
	"fmt"
	"reflect"
)

type ReServer struct {
	M map[string]interface{}
}

// RegisterService 注册服务
func (this *ReServer) RegisterService(service interface{}) {
	serviceType := reflect.TypeOf(service).Elem()
	ServiceName := serviceType.Name()
	if _, ok := this.M[ServiceName]; ok {
		fmt.Println("service has been registered")
	} else {
		this.M[ServiceName] = service
	}
}

//  Start 服务启动
// c 控制器名
// m 方法名
func (this *ReServer) Start(c, m string) {
	for k, v := range this.M {
		// 里面根据业务逻辑执行想要的方法
		if k == c {
			v := reflect.ValueOf(v)
			v.MethodByName(m).Call(nil)
		}
	}
}
