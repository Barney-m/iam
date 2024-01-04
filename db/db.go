package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectMariadb() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.mariadb.user"), viper.GetString("db.mariadb.password"),
		viper.GetString("db.mariadb.host"), viper.GetInt("db.mariadb.port"),
		viper.GetString("db.mariadb.dbname"))
	fmt.Println("Connection String ------> ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
