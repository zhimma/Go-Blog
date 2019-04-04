package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseTime struct {
	CreatedAt time.Time  `json:"created_at" gorm:"default:null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:null"`
	DeletedAt time.Time  `json:"deleted_at" gorm:"default:null"`
}

var DB *gorm.DB

func init() {
	err := initDbConnect()
	if err != nil {
		panic("failed to connect database," + err.Error())
	}
	DB.AutoMigrate(&User{})
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}

func initDbConnect() (err error) {
	dsn := beego.AppConfig.String("dsn")
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
