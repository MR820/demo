package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is jsjxzw@163.com
email1 is abc@df.org
email2 is   kk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	math := re.FindAllStringSubmatch(text, -1)
	fmt.Println(math)
}
