package service

import (
	"fmt"
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /usr/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /usr/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.PassWord = password
	_, err := models.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "创建用户成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /usr/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	Id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(Id)
	_, err := models.DeleteUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /usr/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	Id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(Id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	//fmt.Println("thhis xxxxxx ", c.PostForm("name"))
	//fmt.Println("this is a test:", user.ID, user.Name)
	_, err := models.UpdateUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "修改用户成功",
	})
}
