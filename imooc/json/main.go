/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/12
 * Time 上午9:45
 */

package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct { // tag
	ID         string      `json:"id"`
	Item       []OrderItem `json:"item"`
	TotalPrice float64     `json:"total_price"`
}

func main() {
	unmarshal()
}

func marshal() {
	o := Order{
		ID:         "1234",
		TotalPrice: 20,
		Item: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "interview",
				Price: 10,
			},
		},
	}

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func unmarshal() {
	s := `{"id":"1234","item":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":10}],"total_price":20}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}
