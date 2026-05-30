package model

import "time"

type Department struct {
	ID        int          `gorm:"primaryKey"`
	Name      string       `gorm:"not null"`
	ParentID  *int
	CreatedAt time.Time    `gorm:"autoCreateTime"`

	Children  []Department `gorm:"foreignKey:ParentID"`
	Employees []Employee   `gorm:"foreignKey:DepartmentID"`
}