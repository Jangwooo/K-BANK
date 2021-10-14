package model

import (
	"context"
	_ "database/sql"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dbUser     = os.Getenv("mysql_username")
	dbPassword = os.Getenv("mysql_pwd")
	dbAddress  = os.Getenv("mysql_address")
)

const (
	AccessTokenDB  = iota
	RefreshTokenDB = iota
	MobileTokenDB  = iota
)

func ConnectDB() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/K_BANK", dbUser, dbPassword, dbAddress))
}

func ConnectRedis(db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_address"),
		Password: os.Getenv("redis_pwd"),
		DB:       db,
	})
	_, err := client.Ping(context.Background()).Result()

	return client, err
}
