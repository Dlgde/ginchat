package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
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
		"code":    0, //0成功，-1失败
		"message": "获取成功",
		"data":    data,
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
	// TODO 格式化了解
	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1, //0成功，-1失败
			"message": "用户名已经注册",
			"data":    data,
		})
		return
	}
	if password != repassword {
		c.JSON(200, gin.H{
			"code":    -1, //0成功，-1失败
			"message": "两次密码不一致",
			"data":    data,
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	_, err := models.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    0, //0成功，-1失败
		"message": "创建用户成功",
		"data":    data,
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
		"code":    -1, //0成功，-1失败
		"message": "删除用户成功",
		"data":    user,
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success 200 {string} json{"code","message"}
// @Router /usr/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	Id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(Id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配",
		})
	} else {
		_, err = models.UpdateUser(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"message": "修改用户成功",
		})
	}
}

// FindUserByNameAndPwd
// @Summary 用户登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /usr/getUserList [get]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")

	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0成功，-1失败
			"message": "用户不存在",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //0成功，-1失败
			"message": "密码不正确",
			"data":    data,
		})
		return
	}
	data = models.FindUserByNameAndPwd(name, pwd)

	c.JSON(200, gin.H{
		"code":    0, //0成功，-1失败
		"message": "登录成功",
		"data":    data,
	})
}
