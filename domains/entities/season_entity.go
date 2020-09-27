package entities

import (
	"time"
)

// SeasonEntity struct
type SeasonEntity struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
