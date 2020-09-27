package entities

import "time"

// WindEntity struct
type WindEntity struct {
	ID        int `gorm:"primaryKey"`
	Speed     float32
	Deg       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
