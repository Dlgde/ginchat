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
	Phone         string
	Email         string
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
	result := utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord})
	if result.Error != nil {
		fmt.Println("there is a error when delete")
		return nil, result.Error
	}

	return result, nil
}
