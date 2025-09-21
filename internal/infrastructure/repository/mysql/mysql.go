package mysql

import (
	"fmt"
	"sync"

	"github.com/578223592/awesomeCoupon/internal/infrastructure/handler/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var initOnce sync.Once
var client *gorm.DB

func Init() {
	initOnce.Do(func() {
		mysqlDsn := config.Viper.GetString("mysql.dsn")

		db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect database||mysqldsn:=%s", mysqlDsn))
		} else {
			fmt.Printf("connect database success \n")
		}
		client = db
	})
}

func GetClient() *gorm.DB {
	Init() // 避免忘记初始化
	return client
}
