package entities

import (
	"time"
)

// AccountEntity struct
type AccountEntity struct {
	ID        int `gorm:"primaryKey"`
	UserName  string
	PassWord  string
	APIToken  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
