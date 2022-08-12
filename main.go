package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/wickedknock/goCRM/database"
	"github.com/wickedknock/goCRM/lead"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete(lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"))
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connected to DB")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	//defer database.DBConn.Close()

}
