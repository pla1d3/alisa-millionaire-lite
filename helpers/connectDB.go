package helpers

import (
	"database/sql"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func ConnMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:Fex123WellBox@/db_aml")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	return client
}
