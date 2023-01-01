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

// display a specific user in sqlite3 database
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

// update a specific user in sqlite3 database
func UpdateUser(u *model.User) (err error) {
	cmd := `update users set name =?, email =? where id =?`
	_, err = Db.Exec(cmd,
		u.Name,
		u.Email,
		u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func DeleteUser(u *model.User) (err error) {
	cmd := `delete from users where id =?`
	_, err = Db.Exec(cmd,
		u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user model.User, err error) {
	user = model.User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email =?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func CreateSession(u *model.User) (session model.Session, err error) {
	session = model.Session{}
	cmd1 := `insert into sessions(
		uuid,
    email,
		user_id,
		created_at) values (?, ?, ?, ?)`
	_, err = Db.Exec(cmd1,
		createUUID(),
		u.Email,
		u.ID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions where email = ? and user_id = ?`
	err = Db.QueryRow(cmd2, u.Email, u.ID).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)
	return session, err
}
