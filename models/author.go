package models

import "time"

// author models
type Author struct {
	ID        uint64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
