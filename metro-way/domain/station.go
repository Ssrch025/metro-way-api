package domain

import "gorm.io/gorm"

type Station struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Location string // optional
}
