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

func GetTodos() (todos []model.Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func GetTodosByUser(u *model.User) (todos []model.Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id =?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo model.Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}
