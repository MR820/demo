package main

import (
	"time"


)

func main() {
	//手动埋点
	examples.ExampleNewTracer()
	// 用http框架埋点
	examples.HttpExample()
	time.Sleep(5 * time.Second)
}
