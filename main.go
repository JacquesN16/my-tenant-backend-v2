package main
import (
    "log"
    "github.com/gofiber/fiber/v3"
    "github.com/JacquesN16/my-tenant-backend-v2/tenant"
    "github.com/JacquesN16/my-tenant-backend-v2/db"
)
func status(c *fiber.Ctx){
	return c.SendString("Server is up and running !")
}
func setupRoutes(app *fiber.App) {
	app.Get("/")

	app.Get("api/tenants", tenant.getAllTenants)
	app.Post("api/tenant", tenant.insertTenant)
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
