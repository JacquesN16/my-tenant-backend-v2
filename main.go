package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/template/html/v2"
)

import "my-tenant-backend-v2/tenant"
import "my-tenant-backend-v2/db"

func status(c *fiber.Ctx) error{
	return c.SendString("Server is up and running !")
}
func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
        // Render index template
        return c.Render("index", fiber.Map{
            "Title": "Hello, World!",
        })
    })
	app.Get("api/tenants", tenant.GetAllTenants)
	app.Post("api/tenant", tenant.InsertTenant)
}

func main(){
    // Init

    dbErr := db.InitDB()
    engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
            Views: engine,
            ViewsLayout: "layouts/main",
    })

    app.Static("/","./views/public")

    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowCredentials: false,
    }))

    if dbErr != nil {
		log.Fatal(dbErr)
	}

	// Routes
	setupRoutes(app)
    log.Fatal(app.Listen(":1323"))
}
