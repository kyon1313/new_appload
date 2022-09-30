package main

import (
	"log"
	"os"
	"task_management/database"
	"task_management/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App) {
	app.Post("/user", route.AddUser)
	app.Get("/user", route.GetUsers)
	app.Get("/user/:id", route.GetUser)
	app.Delete("/user/:id", route.DeleteUser)
	app.Put("/user/:id", route.UpdateUser)
	//login
	app.Post("login", route.Log)
	app.Post("/loginhayop", route.Log)
	//Preference
	app.Post("/pref", route.AddPref)
	app.Get("/pref", route.GetPrefs)
	app.Get("/pref/:id", route.Getpref)
	app.Delete("/pref/:id", route.DeletePref)
	app.Put("/pref/:id", route.UpdatePref)
	//team
	app.Post("/team", route.AddTeam)
	app.Get("/team", route.GetTeams)
	app.Get("/team/:id", route.Getteam)
	app.Delete("/team/:id", route.DeleteTeam)
	app.Put("/team/:id", route.UpdateTeam)
	//worksapce
	app.Post("/workspace", route.AddWorkspace)
	app.Get("/workspace", route.GetWorkSpaces)
	app.Get("/workspace/:id", route.GetWorkspace)
	app.Delete("/workspace/:id", route.DeleteWorkspace)
	app.Put("/workspace/:id", route.UpdateWorkSpace)
	//task
	app.Post("/task", route.AddTask)
	app.Get("/task", route.GetTasks)
	app.Get("/task/:id", route.Gettask)
	app.Delete("/task/:id", route.DeleteTask)
	app.Put("/task/:id", route.UpdateTask)

}

func main() {

	app := fiber.New()
	Routes(app)
	database.Migration()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // means all user can access the API or you can just specify the user that can use the system (example: facebook.com) only
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", //header that we only accept
	}))
	app.Use(logger.New())

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
