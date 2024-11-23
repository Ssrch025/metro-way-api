package repositories

import (
	"github.com/Ssrch025/metro-way-api/metro-way/domain"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking *domain.Booking) error
	GetByID(id uint) (*domain.Booking, error)
	GetAll() ([]*domain.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingReposiotry(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *domain.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) GetByID(id uint) (*domain.Booking, error) {
	var booking domain.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) GetAll() ([]*domain.Booking, error) {
	var bookings []*domain.Booking
	if err := r.db.Find(bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
