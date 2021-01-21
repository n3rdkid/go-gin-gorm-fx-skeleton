package models

// User -> Model for User Table
type User struct {
	ID        string    `gorm:"primary_key;not null;unique" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Vehicles  []Vehicle `json:"vehicles`
}

// TableName -> Overriding default table name
func (User) TableName() string {
	return "users"
}
