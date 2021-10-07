package model

import (
	"context"
	_ "database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dbUser     = os.Getenv("mysql_root_id")
	dbPassword = os.Getenv("mysql_root_pw")
	dbAddress  = os.Getenv("mysql_address")
)

func ConnectDB() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/K_BANK", dbUser, dbPassword, dbAddress))
}

func ConnectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6370",
		Password: "",
		DB:       0,
	})
}

func setKeyOnRedis(key, value string) error {
	rdb := ConnectRedis()
	defer rdb.Close()
	ctx := context.Background()

	err := rdb.Set(ctx, key, value, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func getValueOnRedis(key string) (string, error) {
	rdb := ConnectRedis()
	defer rdb.Close()
	ctx := context.Background()

	result := rdb.Get(ctx, key)
	if result == nil {
		return "", fmt.Errorf("not found value via key: %s", key)
	}

	return result.String(), nil
}
