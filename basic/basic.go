package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue() {
	a := 0
	s := ""
	fmt.Printf("%d %q\n", a, s)
}

//欧拉公式
func euler() {
	fmt.Printf("%.3f",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)
}

func main() {
	//fmt.Println("runtime.GOARCH")
	//variableZeroValue()

	//euler()
	//
	//triangle()
	//
	//consts()

	//enums()
	//const filename = "abc.txt"
	//
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("%s\n", contents)
	//}

}
