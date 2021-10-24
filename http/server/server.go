/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/28
 * Time 下午2:20
 */

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello other, %q", html.EscapeString(r.URL.Path))
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}

	//sigterm := make(chan os.Signal, 1)
	//signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	//select {
	//case s := <-sigterm:
	//	log.Println(s)
	//}
}
