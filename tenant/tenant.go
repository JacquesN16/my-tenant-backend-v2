package tenant

import (
	"my-tenant-backend-v2/db"
	"github.com/gofiber/fiber/v2"
)

func GetAllTenants(c *fiber.Ctx) error {
	return c.SendString("All tenants")
}

func InsertTenant(c *fiber.Ctx) error {

	newTenant := new(db.Tenant)

	err:= c.BodyParser(newTenant)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})

		return err
	}

	res, err := db.CreateTenant(newTenant.FirstName,newTenant.LastName,newTenant.Email,newTenant.StartDate,newTenant.Rent,newTenant.Charge,newTenant.EndDate)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})

		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"msg": "Tenant created successfully",
		"data": res,
	})

	return nil
}
