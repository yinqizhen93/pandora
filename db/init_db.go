package db

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"pandora/ent"
)

var Client *ent.Client

func InitDB() *ent.Client {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.username")
	passwd := viper.GetString("database.password")
	database := viper.GetString("database.database")
	maxConnPool := viper.GetInt("database.maxConnPool")
	maxIdleConns := viper.GetInt("database.maxIdleConns")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, passwd, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败:%s", err))
	}
	db.SetMaxOpenConns(maxConnPool)
	db.SetMaxIdleConns(maxIdleConns)
	drv := entsql.OpenDB("mysql", db)
	//cDrv := entcache.NewDriver(drv,
	//	entcache.ContextLevel(),
	//	entcache.TTL(time.Minute),             // 缓存过期时间
	//	entcache.Levels(entcache.NewLRU(128)), // 缓存最大条数
	//) // 添加山缓存
	client := ent.NewClient(ent.Driver(drv))
	Client = client
	return client

	// Run the auto migration tool.
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}
	//return client
}
