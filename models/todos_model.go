package models

import "time"

type Todos struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Task      string    `json:"task"`
	Completed string    `gorm:"default:'pending'"`
	CreatedAt time.Time `json:"created_at"`
}
