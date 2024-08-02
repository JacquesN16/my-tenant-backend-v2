package db

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	StartDate int64   `json:"startDate"`
	Rent      float64 `json:"rent"`
	Charge    float64 `json:"charge"`
	EndDate   *int64   `json:"endDate"`
}

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("tenant.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Tenant{})

	return nil
}

func CreateTenant(firstName string, lastName string, email string, startDate int64, rent float64, charge float64){
	var newTenant = Tenant{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		StartDate: startDate,
		Rent: rent,
		Charge: charge,
		EndDate: nil,
	}

	db, err := gorm.Open(sqlite.Open("tenant.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Create(newTenant)
}
