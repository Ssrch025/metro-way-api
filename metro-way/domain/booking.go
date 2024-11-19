package domain

import "time"

type Booking struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	TicketID      uint      `gorm:"not null"`
	TravellerInfo string    `gorm:"type:json"`
	PaymentStatus string    `gorm:"type:varchar(20);not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type BookingRepository interface {
	CreateTicket(booking *Booking) error
	GetTicketByID(id uint) (*Booking, error)
	GetAllTicket() ([]*Booking, error)
}
