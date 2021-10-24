package model

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type TransformedUser struct {
	ID   uint   `json:"id" uri:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
