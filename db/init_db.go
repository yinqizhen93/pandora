package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pandora/ent"
	"pandora/logs"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.username")
	passwd := viper.GetString("database.password")
	database := viper.GetString("database.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, passwd, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to mysql")
	}
	sqlDb, err := db.DB()
	if err != nil {
		logs.Logger.Error(fmt.Sprintf("数据库初始化失败:%s", err))
		panic(err)
	}
	maxConnPool := viper.GetInt("database.maxConnPool")
	maxIdleConns := viper.GetInt("database.maxIdleConns")
	sqlDb.SetMaxOpenConns(maxConnPool)
	sqlDb.SetMaxIdleConns(maxIdleConns)
	DB = db
	return DB
}

func InitDB2() *ent.Client {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.username")
	passwd := viper.GetString("database.password")
	database := viper.GetString("database.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, passwd, host, port, database)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		panic("fail to connect to mysql")
	}
	//sqlDb, err := db.DB()
	//if err != nil {
	//	service.Logger.Error(fmt.Sprintf("数据库初始化失败:%s", err))
	//	panic(err)
	//}
	//maxConnPool := viper.GetInt("database.maxConnPool")
	//maxIdleConns := viper.GetInt("database.maxIdleConns")
	//sqlDb.SetMaxOpenConns(maxConnPool)
	//sqlDb.SetMaxIdleConns(maxIdleConns)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
