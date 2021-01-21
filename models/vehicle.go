package models

import "gorm.io/gorm"

// Vehicle -> Model for vehicles table
type Vehicle struct {
	gorm.Model
	VehicleType string `json:"vehicle_type"`
	User        User   `gorm:"constraint:onUpdate:CASCADE, onDelete: SET NULL;"`
}

// TableName -> Overriding default table name
func (Vehicle) TableName() string {
	return "vehicles"
}
