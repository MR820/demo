/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/29
 * Time 上午10:50
 */

package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func test1() {
	tty, err := os.OpenFile("/Users/xingzhiwei/go/src/learngo/abc.txt", os.O_RDWR, 0)
	defer tty.Close()
	if err != nil {
		panic(err)
	}
	var r io.Reader
	r = tty
	fmt.Printf("&r:%p,type:%v \n", r, reflect.TypeOf(r))
	var w io.Writer
	w = r.(io.Writer)
	fmt.Printf("&w:%p,type:%v \n", w, reflect.TypeOf(w))
}

func test2() {
	num := 1.2345
	fmt.Println("type:", reflect.TypeOf(num))
	fmt.Println("value:", reflect.ValueOf(num))
}

func test3() {
	num := 1.2345
	// 反射类型对象
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	fmt.Println(reflect.TypeOf(pointer))
	fmt.Println(reflect.TypeOf(value))

	// 接口类型变量
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)

	fmt.Println(reflect.TypeOf(convertPointer))
	fmt.Println(reflect.TypeOf(convertValue))
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is:", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v \n", field.Name, field.Type, value)
	}
}

func test4() {
	// reflect.value 设置实际变量的值
	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()
	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name:", name, ",age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func test5() {
	user := User{1, "Allen.Wu", 25}

	getValue := reflect.ValueOf(user)

	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")

	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	methodValue.Call(args)

	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}

func main() {
	test1()
	test2()
	test3()
	user := User{
		1,
		"Allen.Wu",
		25,
	}
	DoFiledAndMethod(user)
	test4()
	test5()

}
