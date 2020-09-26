package entities

import "time"

// AccountEntity struct
type AccountEntity struct {
	ID        int
	UserName  string
	PassWord  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
