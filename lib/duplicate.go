package lib

import (
	"fmt"
	"log"

	"K-BANK/model"
)

func DuplicateCheck(table string, column string, value string) bool {
	db, err := model.ConnectDB()
	var result string
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	fmt.Println(table, column, value)

	query := fmt.Sprintf("SELECT EXISTS(select %s from %s where %s = ?)", column, table, column)

	_ = db.Get(&result, query, value)

	// 0: 중복 없음, 1: 중복 있음
	if result == "1" {
		return false
	}
	return true
}
