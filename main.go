package main

import (
	"fmt"
	"github.com/KrMrityunjay/go-fiber-crm/lead"

	"github.com/KrMrityunjay/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App){
	app.Get("api/v1/lead",lead.Getleads)
	app.Get("api/v1/lead/:id",lead.GetLead)
	app.Post("api/v1/lead",lead.NewLead)
	app.Delete("api/v1/lead/:id",lead.DeleteLead)

}
func initDatabse(){
	var err error
	database.DBConn,err = gorm.Open("sqlite3","leads.db")
	if err != nil{
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main()  {
	app := fiber.New()
	initDatabse()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
	
}