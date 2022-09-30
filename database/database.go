package database

import (
	"fmt"
	"task_management/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var DSN = "host=ec2-18-208-55-135.compute-1.amazonaws.com user=rtdxajczuwsukk password=ff9ab4e5d2ec82a0b411c7fc8d51d0816711adaf25d76ce56e78710be8682194 dbname=df6tcrni6lh2bk port=5432"

func Migration() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the databse")
	DB.AutoMigrate(&model.User{}, &model.Preference{}, &model.Team{}, &model.Workspace{}, &model.Task{})
}
