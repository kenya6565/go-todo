package main

import (
	"fmt"
	"go-todo/app/controllers"
	"go-todo/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(models.Db)
	user, _ := models.GetUserByEmail("test@test.com")
	fmt.Println(user)
	session, err := models.CreateSession(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(session)
	controllers.StartMainServer()

}
