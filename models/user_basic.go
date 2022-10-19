package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email""`
	Identity      string
	ClientIp      string
	LoginTime     time.Time
	HeartBeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}
func FindUserByPhone(phone string) *gorm.DB {
	user := &UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}
func FindUserByEmail(email string) *gorm.DB {
	user := &UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) (*gorm.DB, error) {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("there is a error when create")
		return nil, result.Error
	}

	return result, nil
}

func DeleteUser(user UserBasic) (*gorm.DB, error) {
	result := utils.DB.Delete(&user)
	if result.Error != nil {
		fmt.Println("there is a error when delete")
		return nil, result.Error
	}

	return result, nil
}

func UpdateUser(user UserBasic) (*gorm.DB, error) {
	result := utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Email: user.Email, Phone: user.Phone})
	if result.Error != nil {
		fmt.Println("there is a error when delete")
		return nil, result.Error
	}

	return result, nil
}
