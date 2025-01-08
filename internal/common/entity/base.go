package entity

import "time"

type BaseEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
