package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

import "my-tenant-backend-v2/tenant"
import "my-tenant-backend-v2/db"

func status(c *fiber.Ctx) error{
	return c.SendString("Server is up and running !")
}
func setupRoutes(app *fiber.App) {
	app.Get("/", status)
	app.Get("api/tenants", tenant.GetAllTenants)
	app.Post("api/tenant", tenant.InsertTenant)
}

func main(){
    // Init
    app := fiber.New()
    dbErr := db.InitDB()

    if dbErr != nil {
		log.Fatal(dbErr)
	}

	// Routes
	setupRoutes(app)
    log.Fatal(app.Listen(":1323"))
}
