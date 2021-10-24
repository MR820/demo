/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/8/28
 * Time 2:09 下午
 */

package user

import (
	"fmt"
	"net/http"

	"learngo/api/mysql"

	"learngo/api/model"

	"github.com/gin-gonic/gin"
)

var db = mysql.DbBms

const (
	deleteUserById    = "delete from user where id = ?"
	updateUserAgeById = "update user set age=? where id=?"
	findUserById      = "select id,age,`name` from user where id = ?"
	findUserList      = "select id,age,`name` from user limit 10"
	createUser        = "INSERT INTO user(id,`name`,age) values(?,?,?)"
)

func DeleteUser(c *gin.Context) {

	var user model.User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	res, err := db.Exec(deleteUserById, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	_, err := db.Exec(updateUserAgeById, user.Age, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "更新成功"})
}

func DetailUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	row := db.QueryRow(findUserById, user.Id)
	err := row.Scan(&user.Id, &user.Age, &user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func ListUser(c *gin.Context) {

	var user model.User
	var list []model.User

	rows, err := db.Query(findUserList)
	defer rows.Close()

	if err != nil {
		fmt.Println("Query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Age, &user.Name); err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		list = append(list, user)
	}

	c.JSON(http.StatusOK, gin.H{"data": list})

}

func InitPage(c *gin.Context) {
	fmt.Println(db)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(createUser, user.Id, user.Name, user.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"msg": "创建成功"})
}
