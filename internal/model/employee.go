package model

import "time"

type Employee struct {
	ID           int    `gorm:"primaryKey"`
	DepartmentID int    `gorm:"not null"`
	FullName     string `gorm:"not null"`
	Position     string `gorm:"not null"`
	HiredAt      *time.Time
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
