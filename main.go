package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/mehfooz-ul-haq/go-fiber-crm/database"
	"github.com/mehfooz-ul-haq/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.NewLead)
	app.Delete("/api/v1/leads/:id", lead.DeleteLead)
}

func initDb() {
	var err error
	database.DbConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Error conecting database")
	}

	fmt.Println("Connnection opened to database")
	database.DbConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDb()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DbConn.Close()
}
