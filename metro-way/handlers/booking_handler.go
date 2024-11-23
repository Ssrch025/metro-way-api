package handlers

import (
	"github.com/Ssrch025/metro-way-api/metro-way/domain"
	"github.com/Ssrch025/metro-way-api/metro-way/usecases"
	"github.com/gofiber/fiber/v2"
)

type BookingHandler struct {
	bookingUsecase usecases.BookingUsecase
}

func NewBookingHandler(app *fiber.App, bookingUsecase usecases.BookingUsecase) {
	handler := &BookingHandler{bookingUsecase: bookingUsecase}

	bookingRoutes := app.Group("/bookings")
	bookingRoutes.Post("/", handler.CreateBooking)
	bookingRoutes.Get("/:id", handler.GetBookingById)
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var booking domain.Booking
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.bookingUsecase.CreateBooking(&booking); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(booking)
}

func (h *BookingHandler) GetBookingById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Invalid booking ID"})
	}

	booking, err := h.bookingUsecase.GetBookingById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(booking)
}
