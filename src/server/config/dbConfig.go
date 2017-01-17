package config

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

func SetUpDBConfig() {
	SetUpShopDBConfig()
}

func SetUpShopDBConfig() {
	dialect := viper.GetString("db.dialect")        //"mysql"
	host := viper.GetString("db.host")        //"192.168.1.204"
	port := viper.GetString("db.port")        //"3306"
	user := viper.GetString("db.user")        //"web"
	password := viper.GetString("db.password")        //"123456"
	database := viper.GetString("db.database")        //"shop"
	maxIdle, _ := strconv.Atoi(viper.GetString("db.maxIdle"))        //"20"
	maxOpen, _ := strconv.Atoi(viper.GetString("db.maxOpen"))        //"20"

	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open(dialect, url)
	if err != nil {
		log.Println("连接shop数据库失败:" + err.Error())
	}else {
		log.Println("成功连接shop数据库...")
	}
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)

	ShopDB = db
}

var (
	ShopDB *sql.DB
)
