/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/6
 * Time 下午1:34
 */

package main

import (
	"fmt"
	"net/http"

	//_ "net/http/pprof"
	_ "github.com/mkevac/debugcharts"
)

type Employee struct {
	name   string
	salary int
	sales  int
	bonus  int
}

const BONUS_PERCENTAGE = 10

func getBonusPercentage(salary int) int {
	percentage := (salary * BONUS_PERCENTAGE) / 100
	return percentage
}

func findEmployeeBonus(salary, noOfSales int) int {
	bonusPercentage := getBonusPercentage(salary)
	bonus := bonusPercentage * noOfSales
	return bonus
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	//http.ListenAndServe(":10107", nil)
	http.ListenAndServe(":10108", nil)
	var john = Employee{"John", 500, 5, 0}
	john.bonus = findEmployeeBonus(john.salary, john.sales)
	fmt.Println(john.bonus)
}
