package models

import (
	"go-todo/model"
	"log"
	"time"
)

// generating users table in sqlite3 database
func CreateUser(u *model.User) (err error) {
	cmd := `insert into users(
		uuid,
		name,
    email,
    password,
		created_at) values (?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user model.User, err error) {
	user = model.User{}
	cmd := `select id, uuid, name, email, password, created_at 
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}
