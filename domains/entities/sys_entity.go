package entities

import "time"

// SysEntity struct
type SysEntity struct {
	ID          int `gorm:"primaryKey"`
	Type        string
	CountryCode string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
