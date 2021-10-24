/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/30
 * Time 9:48 上午
 */

package main

import "net/http"

func main() {
	http.HandleFunc("/", handleReq)
	http.ListenAndServe(":3005", nil)
}

func handleReq(w http.ResponseWriter, r *http.Request) {

}
