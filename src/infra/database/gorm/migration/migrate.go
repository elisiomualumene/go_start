package main

import (
	"go_start/src/app/post/model"
	database "go_start/src/infra/database/gorm"
)

func init(){
	database.ConnectDB()
}

func main(){
	database.DB.AutoMigrate(&model.Post{})
}