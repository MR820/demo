/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/29
 * Time 11:47 上午
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", response)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln("ListenAndServer: ", err)
	}
}
