package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/gin-gorm/dao"
	"imooc.com/ccmouse/learngo/gin-gorm/model"
)

var db = dao.GetDB()

func CreateUser(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	ret := db.Table("user").Create(&user)
	if ret.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": ret.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "创建成功", "data": user.ID})
}

func FetchAllUser(c *gin.Context) {
	var users []model.UserModel
	var _users []model.TransformedUser
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	limit, err := strconv.Atoi(size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	offset := (p - 1) * limit
	ret := db.Table("user").Limit(limit).Offset(offset).Find(&users)
	if ret.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": ret.Error.Error()})
		return
	}
	for _, item := range users {
		_users = append(_users, model.TransformedUser{
			ID:   item.ID,
			Name: item.Name,
			Age:  item.Age,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "查询列表成功", "data": gin.H{"count": ret.RowsAffected, "list": _users}})
}

func FetchSingleUser(c *gin.Context) {
	var user model.UserModel
	var _user model.TransformedUser
	if err := c.ShouldBindUri(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusOK, "msg": err.Error()})
		return
	}
	ret := db.Table("user").First(&user, _user.ID)
	if ret.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": ret.Error.Error()})
		return
	}
	_user = model.TransformedUser{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "查询详情成功", "data": _user})
}

func UpdateUser(c *gin.Context) {
	var user model.UserModel
	var _user model.TransformedUser
	if err := c.ShouldBindUri(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	fmt.Println(_user.ID)
	ret := db.Table("user").Model(&_user).Updates(&user)
	if ret.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": ret.Error.Error()})
		return
	}
	_user = model.TransformedUser{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "更新成功", "data": _user})
}

func DeleteUser(c *gin.Context) {
	//var user model.UserModel
	var _user model.TransformedUser
	if err := c.ShouldBindUri(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ret := db.Table("user").Delete(&_user)
	if ret.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": ret.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "删除成功", "data": ret.RowsAffected})
}
