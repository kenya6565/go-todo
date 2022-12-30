package models

import (
	"go-todo/model"
	"log"
	"time"
)

func CreateTodo(u *model.User, content string) (err error) {
	cmd := `insert into todos(
    content,
		user_id,
    created_at) values (?, ?, ?)`
	_, err = Db.Exec(cmd,
		content,
		u.ID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err

}

// display a specific todo in sqlite3 database
func GetTodo(id int) (todo model.Todo, err error) {
	todo = model.Todo{}
	cmd := `select id, content, user_id, created_at 
	from todos where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}
