package model

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func (u User) LoginWithID(id, pwd string) (bool, error) {
	db, err := ConnectDB()
	if err != nil {
		return false, fmt.Errorf("can't connect to database: %s", err)
	}
	defer db.Close()

	err = db.Get(&u, "SELECT * FROM user where user_id = ?", id)

	match := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if match != nil {
		return false, fmt.Errorf("not match password")
	}
	return true, nil
}
