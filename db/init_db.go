package db

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"pandora/ent"
	"pandora/service/cache"
)

var Client *ent.Client

func InitDB() {
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
	Client = ent.NewClient(ent.Driver(drv))
	// add runtime hooks
	Client.Use(removeCache)
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// updateCache is a hook to remove related cache
func removeCache(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		value, err := next.Mutate(ctx, m)
		if err == nil { // todo 这里的错误会不会有不影响缓存更新的，
			//if m.Op() >= 1 && m.Op() <= 5 {
			// remove cache here
			cache.Cache.DelBySchema(m.Type()) // todo 批量上传会导致缓存一直更新
			//}
		}
		return value, err
	})
}
