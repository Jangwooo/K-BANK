package data

import (
	"database/sql"
)

type User struct {
	UserId      string         `db:"user_id"`
	Password    string         `db:"password"`
	PhoneNumber string         `db:"phone_number"`
	SSN         string         `db:"SSN"`
	Name        string         `db:"name"`
	Nickname    sql.NullString `db:"nickname"`
	Agree       string         `db:"agree"`
}

func (u User) SignIn() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}
