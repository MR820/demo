/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/10/25
 * Time 上午12:19
 */

package logger

import (
	"io"
	"log"
	"os"
)

var Info *log.Logger

var Error *log.Logger

func init() {
	infoFile, err := os.OpenFile("./info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	defer infoFile.Close()
	if err != nil {
		log.Fatalln("打开日志文件失败", err)
	}

	Info = log.New(os.Stdout, "Info", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(os.Stderr, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)

	errorFile, err := os.OpenFile("./error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	defer errorFile.Close()
	if err != nil {
		log.Fatalln("打开日志文件失败", err)
	}

	Error = log.New(os.Stdout, "Info", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(os.Stderr, errorFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
}
