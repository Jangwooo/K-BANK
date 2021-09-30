package data

import (
	_ "database/sql"
	"fmt"
	"os"

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
