package sendMail

import (
	"github.com/gofiber/fiber/v2"
	"my-tenant-backend-v2/db"
)

func SendMail(c *fiber.Ctx) {
	// Get the tenant ID from the URL
	id := c.Params("id")
	println("ID: ", id)
	// Get the tenant from the database
	tenant, err := db.GetTenant(id)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
	}

	// Send the email
	err = sendEmail(tenant)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})
	}


	return c.Render("succes", fiber.Map{})
}


func sendEmail(tenant db.Tenant) error {
	// Send the email
	// define the email body
	body := "Hello " + tenant.FirstName + " " + tenant.LastName + ",\n\n"
	receiver := tenant.Email
	subject := "Rent Payment Reminder"
	// send the email

}
