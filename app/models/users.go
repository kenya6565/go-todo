package models

import (
	"go-todo/model"
	"log"
	"time"
)

// generating users table in sqlite3 database
func (u *model.User) CreateUser() (err error) {
	cmd := `insert into users(
		uuid,
		name,
    email,
    password,
		created_at) values (?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd,
		createUUID,
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
