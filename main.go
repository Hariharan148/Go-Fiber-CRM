package main

import (
	"fmt"

	"github.com/Hariharan148/Go-Fiber-CRM/database"
	"github.com/Hariharan148/Go-Fiber-CRM/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)


func SetUpRoutes(app *fiber.App){
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func InitDb(){
	var err error
	database.DbCon, err = gorm.Open("sqlite3", "lead.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DbCon.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}




func main(){
	app := fiber.New()
	InitDb()
	SetUpRoutes(app)
	app.Listen(3000)
	defer database.DbCon.Close()
}