package domain

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	UserID   uint `gorm:"not null"`
	TicketID uint `gorm:"not null"`
	// TravellerInfo string `gorm:"type:json"`                 // JSON format
	PaymentStatus string `gorm:"type:varchar(20);not null"` // Pending, Completed, Failed
}
