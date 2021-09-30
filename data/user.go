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
	UserType    string         `db:"user_type"`
}

func (u User) SignIn() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if u.Nickname.Valid {
		_, err = db.NamedExec("INSERT INTO user values (:user_id, :password, :phone_number, :SSN, :name, :nickname, :agree, :user_type)", &u)
	} else {
		_, err = db.NamedExec("INSERT INTO user(user_id, password, phone_number, SSN, name, agree) values (:user_id, :password, :phone_number, :SSN, :name, :agree)", &u)
	}

	if err != nil {
		return err
	}
	return nil
}
