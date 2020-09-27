package entities

import "time"

// UserEntity struct
type UserEntity struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	FirstName string
	LastName  string
	Avatar    string
	Bio       string
	RoleID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
