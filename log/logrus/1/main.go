/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/17
 * Time 上午10:48
 */

package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Info("A walrus appears")
}
