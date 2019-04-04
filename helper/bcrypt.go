package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (encryptPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("加密失败")
	}

	encryptPassword = string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encryptPassword, nil
}

func CheckEncrypt(encodeValue, checkValue string) (status bool) {
	err := bcrypt.CompareHashAndPassword([]byte(encodeValue), []byte(checkValue))
	if err != nil {
		return false
	}
	return true
}
