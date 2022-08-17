package db

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"log"
	"pandora/ent"
	"pandora/ent/migrate"
	"pandora/service/cache"
	"pandora/service/config"
)

func NewEntClient(cache cache.Cacher, conf config.Config) *ent.Client {
	maxConnPool := conf.GetInt("database.maxConnPool")
	maxIdleConns := conf.GetInt("database.maxIdleConns")
	dsn := getDsn(conf)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败:%s", err))
	}
	db.SetMaxOpenConns(maxConnPool)
	db.SetMaxIdleConns(maxIdleConns)
	drv := entsql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv))
	// add runtime hooks
	client.Use(removeCache(cache))
	// Run the auto migration tool.
	log.Println("start to migrate schema..")
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("successful migrate schema..")
	return client
}

func getDsn(conf config.Config) string {
	host := conf.GetString("database.host")
	port := conf.GetString("database.port")
	user := conf.GetString("database.username")
	passwd := conf.GetString("database.password")
	database := conf.GetString("database.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, passwd, host, port, database)
	return dsn
}

// updateCache is a hook to remove related cache
func removeCache(cache cache.Cacher) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			if err == nil { // todo 这里的错误会不会有不影响缓存更新的，
				//if m.Op() >= 1 && m.Op() <= 5 {
				// remove cache here
				cache.DelBySchema(m.Type()) // todo 批量上传会导致缓存一直更新, 高并发时如何控制
				//}
			}
			return value, err
		})
	}
}

var ProviderSet = wire.NewSet(NewEntClient)
