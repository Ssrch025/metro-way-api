package usecases

import (
	"github.com/Ssrch025/metro-way-api/metro-way/domain"
	"github.com/Ssrch025/metro-way-api/metro-way/repositories"
)

type BookingUsecase interface {
	CreateBooking(booking *domain.Booking) error
	GetBookingById(id uint) (*domain.Booking, error)
}

type bookingUsecase struct {
	bookingRepo repositories.BookingRepository
}

func NewBookingUsecase(bookingRepo repositories.BookingRepository) BookingUsecase {
	return &bookingUsecase{bookingRepo: bookingRepo}
}

func (u *bookingUsecase) CreateBooking(booking *domain.Booking) error {
	// business logic
	return u.bookingRepo.Create(booking)
}

func (u *bookingUsecase) GetBookingById(id uint) (*domain.Booking, error) {
	// business logic
	return u.bookingRepo.GetByID(id)
}
