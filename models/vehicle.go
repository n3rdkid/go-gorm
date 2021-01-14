package models

import "gorm.io/gorm"

// Vehicle -> Vehicle users
type Vehicle struct {
	gorm.Model
	TypeOfVehicle string
	NoOfWheels    int
}
