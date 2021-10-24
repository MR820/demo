/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/10/14
 * Time 上午10:45
 */

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	walkDir("/Users/xingzhiwei/go/src/learngo")
}

func walkDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filename := file.Name()
		filepath.Walk(filename, func(path string, fi os.FileInfo, err error) error {
			depth := strings.Count(path, "/") - strings.Count(filename, "/")
			if depth > depth {
				return filepath.SkipDir
			}
			if err != nil {
				//处理文件读取异常
			}
			if fi.IsDir() {
				//满足条件不用管
				//不满足条件
				return filepath.SkipDir
			}
			return nil
		})
	}
}
