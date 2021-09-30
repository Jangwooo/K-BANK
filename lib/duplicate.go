package lib

import (
	"fmt"
	"log"

	"K-BANK/data"
)

func DuplicateCheck(table string, column string, value string) bool {
	db, err := data.ConnectDB()
	var result string
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	fmt.Println(table, column, value)

	query := fmt.Sprintf("select exists(select * from %s where %s = ?)", table, column)

	err = db.QueryRow(query, value).Scan(&result)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)

	// 0: 중복 없음, 1: 중복 있음
	if result == "1" {
		return false
	}
	return true
}
