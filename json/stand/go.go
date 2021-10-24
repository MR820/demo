/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/17
 * Time 上午11:36
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

func main() {
	dir, _ := os.Getwd()
	fileName := path.Join(dir, "20210917.log")
	fd, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err", err)
	}
	defer fd.Close()
	br := bufio.NewReader(fd)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		var l log
		json.Unmarshal(a, &l)
		fmt.Println(l.Time)
		s := l.Time.Format("2006-01-02 15:04:05")
		fmt.Println(s)
	}
}

type log struct {
	Animal string    `json:"animal"`
	Level  string    `json:"level"`
	Msg    string    `json:"msg"`
	Time   time.Time `json:"time"`
	Number *int      `json:"number,omitempty"`
	Size   *int      `json:"size,omitempty"`
}
