package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(10.19.106.217:3306)/yundabms")
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world|\n")
}

func ArticleDetailServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	rows, err := db.Query("select uid,username from user where uid = ?", id)
	if err != nil {
		panic(err)
	}
	v := reflect.ValueOf(rows)
	fmt.Println(v)
	printResult(rows)
	io.WriteString(w, "hello user\n")
}

func printResult(query *sql.Rows) {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	for k, v := range results { //查询出来的数组
		fmt.Println(k, v)
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", HelloServer)
	router.HandleFunc("/user/{id}", ArticleDetailServer)
	log.Fatal(http.ListenAndServe(":12345", router))
}
