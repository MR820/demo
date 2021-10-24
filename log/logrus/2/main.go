/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/17
 * Time 上午10:50
 */

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func main() {
	dir, _ := os.Getwd()
	fileName := path.Join(dir, "20210917.log")
	if !isFile(fileName) {
		os.Create(fileName)
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("err", err)
	}
	defer src.Close()

	var log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)

	//log.Formatter = new(logrus.TextFormatter)
	//log.Formatter.(*logrus.TextFormatter).DisableColors = true // remove colors
	//log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	log.Level = logrus.TraceLevel
	log.Out = src

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			log.WithFields(logrus.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!")
		}
	}()
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 0,
	}).Trace("Went to the beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

// 文件或文件夹是否存在
func FileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil // 不存在
	}
	return false, err
}

func isDir(dir string) bool {
	info, err := os.Stat(dir)
	if err == nil {
		return false
	}
	return info.IsDir()
}

func isFile(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
