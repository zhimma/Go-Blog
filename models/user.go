package models

import (
	"Blog/helper"
)

type User struct {
	Id       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Phone    string `json:"phone" gorm:"type:varchar(16);not null;default:'';unique_index"`
	Nickname string `json:"nickname" gorm:"type:varchar(40);not null;default:''"`
	Account  string `json:"account" gorm:"type:varchar(12);not null;default:'';unique_index"`
	Password string `json:"password" gorm:"type:varchar(64);not null;default:''"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);not null;default:''"`
	Status   int    `json:"status" gorm:"type:tinyint(1);not null;default:0"`
	Token    string `json:"token" gorm:"-"`
	BaseTime
}

//返回带前缀的表名
func (u *User) TableName() string {
	return TableName("user")
}

func (u *User) CheckAuth() (bool bool, message string) {
	password := u.Password
	err := DB.Debug().Where("account = ?", u.Account).Find(&u).Error
	if err != nil {
		return false, err.Error()
	}
	res := helper.CheckEncrypt(u.Password, password)
	if res == false {
		return false, "用户账号密码错误"
	}
	if u.Status == 1 {
		return false, "用户已被禁用"
	}
	return true, "校验成功"
}

func (u *User) CreateUser() (err error) {
	err = DB.Create(&u).Error
	return
}

func (u *User) UserList(page, size int) (user []User, err error) {
	return user, DB.Debug().Offset((page - 1) * size).Limit(size).Find(&user).Error
}

func (u *User) UserCount() (count int, err error) {
	return count, DB.Model(&User{}).Count(&count).Error
}

func (u *User) GetUserDetailByFiled(key string, value interface{}) (err error) {
	return DB.Debug().Where(key+" = ?", value).Find(&u).Error
}

func (u *User) DeleteUserDetailByFiled(key string, value interface{}) (err error) {
	return DB.Debug().Where(key+" = ?", value).Delete(&u).Error
}

func (u *User) UpdateUserInfo(key string, value interface{}) (err error) {
	return DB.Model(&u).Debug().Where(key+" = ?", value).Update(&u).Error
}
