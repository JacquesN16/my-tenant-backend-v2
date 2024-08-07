package tenant

import (
	"time"
	"strconv"
	"my-tenant-backend-v2/db"
	"github.com/gofiber/fiber/v2"
)

type TenantViewData struct{
	ID int64
	Name string
	StartDate string
	Rent float64
	Charge float64
	Total float64
	Months int32
}

const (
	DDMMYYYY= "02-01-2006"
)

func GetTenantById(c *fiber.Ctx) (TenantViewData, error) {
	id := c.Params("id")
	idInt, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})

	}
	tenant, err := db.GetTenant(idInt)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"msg": err,
			"data":nil,
		})

	}

	var tenantViewData TenantViewData

	tenantViewData.Name = tenant.FirstName + " " + tenant.LastName

   	unixTS := int64(tenant.StartDate)
    tm := time.Unix(unixTS/1000, 0)

    tenantViewData.ID = tenant.ID
    tenantViewData.StartDate = tm.Format(DDMMYYYY)
    tenantViewData.Rent = tenant.Rent
    tenantViewData.Charge = tenant.Charge
    tenantViewData.Total = tenant.Rent + tenant.Charge
    tenantViewData.Months = int32(time.Since(tm).Hours()/720)

    return tenantViewData, nil
}


func GetTenants(c *fiber.Ctx) ([]TenantViewData, error) {
        res, err := db.GetAllTenants()
        if err != nil {
	        c.Status(400).JSON(&fiber.Map{
				"success": false,
				"msg": err,
				"data":nil,
			})
			return nil, err
        }

        var tenants []TenantViewData
        for _, tenant := range res {
        	var tenantViewData TenantViewData
         	tenantViewData.Name = tenant.FirstName + " " + tenant.LastName

          	unixTS := int64(tenant.StartDate)
            tm := time.Unix(unixTS/1000, 0)

            tenantViewData.ID = tenant.ID
            tenantViewData.StartDate = tm.Format(DDMMYYYY)
            tenantViewData.Rent = tenant.Rent
            tenantViewData.Charge = tenant.Charge
            tenantViewData.Total = tenant.Rent + tenant.Charge
            tenantViewData.Months = int32(time.Since(tm).Hours()/720)

            tenants = append(tenants, tenantViewData)
        }


        return tenants, nil
}

func GetAllTenants(c *fiber.Ctx) error {

	res, err := db.GetAllTenants()
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
		"msg": "Tenants fetched successfully",
		"data": res,
	})

	return nil
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
